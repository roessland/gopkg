package mathutil

func AbsInt64(n int64) int64 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func AbsInt(n int) int {
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

func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func SignInt64(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	return 1
}

func SignInt(n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	return 1
}