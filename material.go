package yogort

import "image/color"

type Material struct {
	Color   color.RGBA
	Diffuse float64
}

type Color struct {
	R, G, B uint8
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	a |= 255 << 8
	return
}
