package optimize

import "math"

// FindRootNewtons1D finds a root of the real-valued function
// f(x). That is, x such that f(x) ∈ [-tol, tol]. f must be a continuous and
// differentiable function that can be locally approximated with a
// straight line tangent.
func FindRootNewtons1D(
	f func(x float64) float64,
	dfdx func(x float64) float64,
	x0 float64,
	tol float64,
) float64 {
	x := x0
	fx := f(x)
	for math.Abs(fx) > tol {
		x = x - f(x)/dfdx(x)
		fx = f(x)
	}
	return x
}
