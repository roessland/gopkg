package geom

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Scale(alpha float64) Vec2 {
	return Vec2{
		X: v.X * alpha,
		Y: v.Y * alpha,
	}
}
