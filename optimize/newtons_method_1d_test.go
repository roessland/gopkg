package optimize_test

import (
	"github.com/roessland/gopkg/optimize"
	"github.com/stretchr/testify/require"
	"testing"
)

// Ideal case. Minimum is quickly found.
//
//	f(5.333333) = 0.111111
//	f(5.316667) = 0.000278
//	f(5.316625) = 0.000000
func TestFindRootNewtons1D_ConvergesQuickly(t *testing.T) {
	f := func(x float64) float64 {
		return x*x - 4*x - 7
	}
	dfdx := func(x float64) float64 {
		return 2*x - 4
	}
	x0 := 5.0
	tol := 0.0001

	xMin := optimize.FindRootNewtons1D(f, dfdx, x0, tol, nil)
	require.InDelta(t, 5.31662, xMin, tol)
	require.InDelta(t, 0.0, f(xMin), tol)
}

// Uses 10-20 iterations to find the root the polynomial.
// xk bounces back and forth and it is only by luck that
// we eventually find a root.
//
//	f(0.333333) = 1.000000
//	f(0.166667) = 0.625000
//	f(1.000000) = 25.000000
//	f(0.679487) = 7.432010
//	f(0.463427) = 2.296973
//	f(0.303871) = 0.845971
//	f(0.115009) = 0.696045
//	f(0.475917) = 2.482679
//	f(0.314139) = 0.894592
//	f(0.134983) = 0.661456
//	f(0.568968) = 4.266202
//	f(0.385252) = 1.388072
//	f(0.231397) = 0.640341
//	f(-0.247498) = 1.333159
//	f(-0.927099) = -17.733717
//	f(-0.660908) = -4.811763
//	f(-0.512309) = -1.093519
//	f(-0.452421) = -0.143034
//	f(-0.441888) = -0.004034
//	f(-0.441573) = -0.000004
func TestFindRootNewtons1D_BadStartingPoint(t *testing.T) {
	f := func(x float64) float64 {
		return 27*x*x*x - 3*x + 1
	}
	dfdx := func(x float64) float64 {
		return 3*27*x*x - 3
	}
	x0 := 0.0
	tol := 0.0001

	xMin := optimize.FindRootNewtons1D(f, dfdx, x0, tol, nil)
	require.InDelta(t, -0.44157265, xMin, tol)
	require.InDelta(t, 0.0, f(xMin), tol)
}

// x*2 + 1 has no root, so we will reach max iterations and panic.
func TestFindRootNewtons1D_NoRootMaxIterPanic(t *testing.T) {
	f := func(x float64) float64 {
		return x*2 + 1
	}
	dfdx := func(x float64) float64 {
		return 2 * x
	}
	x0 := 5.0
	tol := 0.0001

	require.Panics(t, func() { optimize.FindRootNewtons1D(f, dfdx, x0, tol, nil) })
}
