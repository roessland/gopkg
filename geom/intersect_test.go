package geom_test

import (
	"github.com/roessland/gopkg/geom"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestIntersectLineCircle(t *testing.T) {
	ps := geom.IntersectLineCircle(geom.Point2{X: 0.0, Y: 0.0}, geom.Point2{X: 1.0, Y: 1.0}, 1)
	assert.InDelta(t, math.Sqrt(2)/2, ps[0].X, 0.0001)
	assert.InDelta(t, math.Sqrt(2)/2, ps[0].Y, 0.0001)
	assert.InDelta(t, -math.Sqrt(2)/2, ps[1].X, 0.0001)
	assert.InDelta(t, -math.Sqrt(2)/2, ps[1].Y, 0.0001)
}

func TestIntersectLineLine(t *testing.T) {
	var l1, l2 geom.Line2
	l1 = geom.Line2{
		P1: geom.Point2{0.0, 0.0},
		P2: geom.Point2{X: 1.0, Y: 1.0},
	}
	l2 = geom.Line2{
		P1: geom.Point2{0.0, 1.0},
		P2: geom.Point2{X: 1.0, Y: 1.0},
	}
	ps := geom.IntersectLineLine(l1, l2)
	assert.InDelta(t, 1.0, ps[0].X, 0.0001)
	assert.InDelta(t, 1.0, ps[0].Y, 0.0001)
}
