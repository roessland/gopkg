package geom

import (
	"math"
)

// IntersectLineCircle finds the points where a line intersects with a circle.
// https://mathworld.wolfram.com/Circle-LineIntersection.html
//
// Line goes through a and b.
// Circle is centered in (0,0) with radius r.
func IntersectLineCircle(p1, p2 Point2, r float64) []Point2 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dr := math.Sqrt(dx*dx + dy*dy)
	D := p1.X*p2.Y - p2.X*p1.Y
	delta := r*r*dr*dr - D*D
	sqrtDelta := math.Sqrt(delta)
	if delta < 0 {
		return nil
	}
	x1 := (D*dy + math.Copysign(1.0, dy)*sqrtDelta) / (dr * dr)
	y1 := (-D*dx + math.Abs(dy)*dx*sqrtDelta) / (dr * dr)
	x2 := (D*dy - math.Copysign(1.0, dy)*sqrtDelta) / (dr * dr)
	y2 := (-D*dx - math.Abs(dy)*dx*sqrtDelta) / (dr * dr)
	return []Point2{
		{x1, y1},
		{x2, y2},
	}
}
