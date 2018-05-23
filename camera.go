package yogort

import (
	"image"
	"image/color"
	"math"
)

type Camera struct {
	Position Vec3
	//LookAt   Vec3
}

type projPlane struct {

}

//func (c Camera) Direction() Vec3 {
//	return c.LookAt.Subtract(c.Position).Normalized()
//}

var red = color.RGBA{255, 0, 0, 255}
var black = color.RGBA{0, 0, 0, 255}

func (c Camera) Render(scene Scene, frameBuffer *image.RGBA) {

	bounds := frameBuffer.Bounds()

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			ray := c.shootRay(bounds, x, y)

			if obj, hit, ok := getFirstIntersection(ray, scene.Objects); ok {
				color := c.CalculateLighting(obj, hit, scene.Lights)
				frameBuffer.SetRGBA(x, y, color)
			} else {
				frameBuffer.SetRGBA(x, y, black)
			}
		}
	}
}

func (c Camera) CalculateLighting(obj Object, pos Vec3, lights []Light) color.RGBA {
	diffuse := 0.9
	specular := 0.3
	n := 20.0

	for _, light := range lights {
		N := obj.NormalAt(pos)
		L := light.Position.Subtract(pos).Normalized()
		R := Reflect(L.Invert(), N)
		V := c.Position.Subtract(pos).Normalized()

		Idiffuse := math.Max(0, diffuse * light.Brightness * Dot(N, L))
		Ispecular := math.Max(0, specular * light.Brightness * math.Pow(Dot(R, V), n))
		return color.RGBA{uint8(255 * (Idiffuse + Ispecular)), 0, 0, 255}
	}

	return red
}

func (c Camera) shootRay(bounds image.Rectangle, x, y int) Ray {
	// TODO remove hardcoded view
	relX := float64(x) / float64(bounds.Max.X)
	relY := float64(y) / float64(bounds.Max.Y)
	posX := lerpFloat(0, 8, relX)
	posY := lerpFloat(0, 6, relY)
	pos := Vec3{posX, posY, 0}
	rayDir := pos.Subtract(c.Position).Normalized()
	return Ray{Origin:c.Position, Direction:rayDir}
}

func lerpFloat(a, b float64, t float64) float64 {
	return (1 - t) * a + t * b
}

func getFirstIntersection(ray Ray, objects []Object) (Object, Vec3, bool) {

	// TODO return real object and calculate real intersection
	for _, o := range objects {
		if hit, ok := o.Intersect(ray); ok {
			return o, hit, true
		}
	}

	return nil, Vec3Zero, false
}
