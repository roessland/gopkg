package iterutil

import (
	"github.com/roessland/gopkg/mathutil"
	"github.com/roessland/gopkg/sliceutil"
	"math"
)

// Return all the elements of the cartesian p-th power of Z_n
//
// For example, Z_3 = {0, 1, 2}, so (Z_3)**2 is the set of ordered pairs
// {  (0,0), (0,1), (0,2), (1,0), (1,1), (1,2), (2,0), (2,1), (2,2)  }.
//
// The result has length n**power.
func CartesianPower(n int64, power int64) <-chan []int64 {
	if power == 0 {
		out := make(chan []int64, 0)
		close(out)
		return out
	}
	if n <= 0 {
		panic("n must be positive")
	}
	if power <= 0 {
		panic("power must be positive")
	}

	// Number of elements in the generated set
	cardinality := int64(math.Pow(float64(n), float64(power)))

	out := make(chan []int64)
	go func() {
		if n == 1 {
			out <- make([]int64, power)
		} else {
			for i := int64(0); i < cardinality; i = i + 1 {
				out <- sliceutil.ZeroPadLeftInt64(mathutil.ToDigits(i, n), power)
			}
		}
		close(out)
	}()
	return out
}

// Return all subsets of Z_n with length k
// There are n choose k such subsets.
// Grossly inefficient, make better algorithm soon.
func Subsets(n int64, length int64) <-chan []int64 {
	out := make(chan []int64)
	go func() {
		for product := range CartesianPower(n, length) {
			if sliceutil.StrictlyIncreasingInt64(product) {
				out <- product
			}
		}
		close(out)
	}()
	return out
}
