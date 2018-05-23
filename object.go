package yogort

import (
	"math"
)

type Object interface {
	Intersect(r Ray) (Vec3, bool)
	NormalAt(at Vec3) Vec3
}

type Sphere struct {
	Position Vec3
	Radius   float64
}

// Line-Sphere intersection -- based on http://www.lighthouse3d.com/tutorials/maths/ray-sphere-intersection/
func (s Sphere) Intersect(r Ray) (Vec3, bool) {
	rayToSphere := s.Position.Subtract(r.Origin)

	if Dot(rayToSphere.Normalized(), r.Direction.Normalized()) < 0 {
		// Sphere is behind ray --> ignore
		return Vec3Zero, false
	}

	pc := r.ProjectPoint(s.Position)
	distToRay := Distance(pc, s.Position)

	if distToRay > s.Radius {
		return Vec3Zero, false
	}

	distPCI := math.Sqrt(s.Radius * s.Radius - distToRay * distToRay)
	distI := Distance(pc, r.Origin) - distPCI

	return r.Origin.Add(r.Direction.Scale(distI)), true
}

func (s Sphere) NormalAt(v Vec3) Vec3 {
	return v.Subtract(s.Position).Normalized()
}

type Plane struct {
	Position Vec3
	Normal   Vec3
	material Material
}

// Line-Plane intersection -- http://geomalgorithms.com/a05-_intersect-1.html
func (p Plane) Intersect(r Ray) (Vec3, bool) {
	w := r.Origin.Subtract(p.Position)
	D := Dot(p.Normal, r.Direction)
	N := -Dot(p.Normal, w)

	if math.Abs(D) < 0.0001 {
		return Vec3Zero, false
	}

	sI := N / D
	// if sI > 1 --> farther away than ray length/dir / maybe not so bad ???
	if sI < 0 {
		//intersection is "behind" us
		return Vec3Zero, false
	}

	return r.Origin.Add(r.Direction.Scale(sI)), true
}

func (p Plane) NormalAt(v Vec3) Vec3 {
	return p.Normal
}
