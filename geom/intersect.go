package geom

import (
	"github.com/roessland/gopkg/mat2"
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

func IntersectLineLine(a Line2, b Line2) []Point2 {
	// a.P1.X + (a.P2.X - a.P1.X)*alpha = b.P1.X + (b.P2.X - b.P1.X)*beta
	// a.P1.Y + (a.P2.Y - a.P1.Y)*alpha = b.P1.Y + (b.P2.Y - b.P1.Y)*beta

	// (a.P2.X - a.P1.X)*alpha + (b.P1.X - b.P2.X )*beta = b.P1.X - a.P1.X
	// (a.P2.Y - a.P1.Y)*alpha + (b.P1.Y - b.P2.Y )*beta = b.P1.Y - a.P1.Y
	A := mat2.Mat2{
		A: a.P2.X - a.P1.X,
		B: b.P1.X - b.P2.X,
		C: a.P2.Y - a.P1.Y,
		D: b.P1.Y - b.P2.Y,
	}
	B := mat2.Vec2{
		B1: b.P1.X - a.P1.X,
		B2: b.P1.Y - a.P1.Y,
	}
	X := A.Solve(B)

	return []Point2{
		a.P1.Add(a.Dir().Scale(X.B1)),
	}
}
