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
