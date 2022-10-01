package geom

type Line2 struct {
	P1, P2 Point2
}

func (l Line2) Dir() Vec2 {
	return l.P2.Sub(l.P1)
}
