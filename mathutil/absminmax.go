package mathutil

func AbsInt64(n int64) int64 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}
