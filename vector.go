package yogort

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

var Vec3Zero = Vec3{0, 0, 0}

func (v Vec3) Normalized() Vec3 {
	len := v.Length()
	return Vec3{
		v.X / len,
		v.Y / len,
		v.Z / len,
	}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{
		v.X + w.X,
		v.Y + w.Y,
		v.Z + w.Z,
	}
}

func (v Vec3) Subtract(w Vec3) Vec3 {
	return Vec3{
		v.X - w.X,
		v.Y - w.Y,
		v.Z - w.Z,
	}
}

func (v Vec3) Scale(s float64) Vec3 {
	return Vec3{
		v.X * s,
		v.Y * s,
		v.Z * s,
	}
}

func (v Vec3) ScaleVec(w Vec3) Vec3 {
	return Vec3{
		v.X * w.X,
		v.Y * w.Y,
		v.Z * w.Z,
	}
}

func (v Vec3) Invert() Vec3 {
	return v.Scale(-1)
}

func Distance(v, w Vec3) float64 {
	return math.Abs(v.Subtract(w).Length())
}

func Dot(v, w Vec3) float64 {
	return v.X * w.X + v.Y * w.Y + v.Z * w.Z
}

func Cross(v, w Vec3) Vec3 {
	return Vec3{
		v.Y * w.Z - v.Z * w.Y,
		v.Z * w.X - v.X * w.Z,
		v.X * w.Y - v.Y * w.X,
	}
}

func Reflect(v, normal Vec3) Vec3 {
	return v.Subtract(normal.Scale(2).Scale(Dot(v, normal)))
}

func Lerp(v, w Vec3, t float64) Vec3 {
	return v.Scale(1 - t).Add(w.Scale(t))
}
