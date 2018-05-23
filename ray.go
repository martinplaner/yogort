package yogort

type Ray struct {
	Origin Vec3
	Direction Vec3
}

// DistToPoint calculates to shortest distance from r to v -- http://geomalgorithms.com/a02-_lines.html
func (r Ray) DistToPoint(p Vec3) float64 {
	Pb := r.ProjectPoint(p)
	return Distance(p, Pb);
}

func (r Ray) ProjectPoint(p Vec3) Vec3 {
	v := r.Direction
	w := p.Subtract(r.Origin)


	c1 := Dot(w,v);
	c2 := Dot(v,v);
	b := c1 / c2;

	return r.Origin.Add(v.Scale(b))
}
