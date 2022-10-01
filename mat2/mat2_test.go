package mat2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const tol = 0.0000001

func TestMat2_Solve_Identity(t *testing.T) {
	A := Mat2{1, 0, 0, 1}
	b := Vec2{5, 6}
	x := A.Solve(b)
	assert.InDelta(t, 5.0, x.B1, tol)
	assert.InDelta(t, 6.0, x.B2, tol)
}

func TestMat2_Solve_Simple1(t *testing.T) {
	A := Mat2{1, 2, 3, 4}
	b := Vec2{1, 2}
	x := A.Solve(b)
	assert.InDelta(t, 0.0, x.B1, tol)
	assert.InDelta(t, 0.5, x.B2, tol)
}

func TestMat2_Solve_Simple2(t *testing.T) {
	// 1 2 @ 1 = 1*1 + 2*2 = 5
	// 3 4   2   3*1 + 4*2   11
	A := Mat2{1, 2, 3, 4}
	b := Vec2{5, 11}
	x := A.Solve(b)
	assert.InDelta(t, 1.0, x.B1, tol)
	assert.InDelta(t, 2.0, x.B2, tol)
}
