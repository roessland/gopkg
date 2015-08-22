package iterutil

import (
    "gitub.com/roessland/gopkg/sliceutil"
)

// Return all the elements of the cartesian p-th power of Z_n
//
// For example, Z_3 = {0, 1, 2}, so (Z_3)**2 is the set of ordered pairs
// {  (0,0), (0,1), (0,2), (1,0), (1,1), (1,2), (2,0), (2,1), (2,2)  }.
//
// The result has length n**power.
func CartesianPower(n int, power int) <-chan []int {
    // Number of elements in the generated set
    cardinality := int(math.Pow(float64(n), float64(power)))

	out := make(chan []int)
	go func() {
		for i := 0; i < cardinality; i = i + 1 {
			out <- sliceutil.ZeroPadLeftInt64(mathutil.ToDigits(i, n), power, 0)
		}
		close(out)
	}()
	return out
}
