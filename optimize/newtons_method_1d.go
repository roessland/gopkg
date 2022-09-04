package optimize

import (
	"fmt"
	"math"
)

type FindRootNewtons1DOpts struct {
	Debug   bool
	MaxIter int
}

// FindRootNewtons1D finds a root of the real-valued function
// f(x). That is, x such that f(x) ∈ [-tol, tol]. f must be a continuous and
// differentiable function that can be locally approximated with a
// straight line tangent.
//
// Example:
func FindRootNewtons1D(
	f func(x float64) float64,
	fprime func(x float64) float64,
	x0 float64,
	tol float64,
	opts *FindRootNewtons1DOpts,
) float64 {
	// Process options
	maxIter := 50
	debug := false
	if opts != nil {
		maxIter = opts.MaxIter
		debug = opts.Debug
	}

	iterationsLeft := maxIter
	x := x0
	fx := f(x)
	for math.Abs(fx) > tol {
		x = x - f(x)/fprime(x)
		fx = f(x)
		iterationsLeft--

		if iterationsLeft == 0 {
			msg := fmt.Errorf("failed to converge after %d iterations", maxIter)
			panic(msg)
		}

		if debug {
			fmt.Printf("f(%f) = %f\n", x, fx)
		}
	}
	return x
}
