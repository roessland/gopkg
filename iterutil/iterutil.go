package iterutil

import (
    "math"
    "github.com/roessland/gopkg/mathutil"
    "github.com/roessland/gopkg/sliceutil"
)

// Return all the elements of the cartesian p-th power of Z_n
//
// For example, Z_3 = {0, 1, 2}, so (Z_3)**2 is the set of ordered pairs
// {  (0,0), (0,1), (0,2), (1,0), (1,1), (1,2), (2,0), (2,1), (2,2)  }.
//
// The result has length n**power.
func CartesianPower(n int64, power int64) <-chan []int64 {
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
