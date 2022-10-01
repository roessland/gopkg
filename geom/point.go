package geom

type Point2 struct {
	X, Y float64
}

func (p1 Point2) Sub(p2 Point2) Vec2 {
	return Vec2{
		X: p2.X - p1.X,
		Y: p2.Y - p1.Y,
	}
}

func (p1 Point2) Add(v Vec2) Point2 {
	return Point2{
		X: p1.X + v.X,
		Y: p1.Y + v.Y,
	}
}
