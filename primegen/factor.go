package primegen

import "github.com/roessland/gopkg/mathutil"

func PollardRho(n int64) int64 {
	var x int64 = 2
	var y int64 = (x*x + 1) % n
	var d int64 = 1

	for d == 1 {
		x = (x*x + 1) % n
		gy := (y*y + 1) % n
		y = (gy*gy + 1) % n
		d = mathutil.GCD(mathutil.AbsInt64(x-y), n)
	}

	if d == n {
		panic("pollard rho failure")
	}
	return d
}
