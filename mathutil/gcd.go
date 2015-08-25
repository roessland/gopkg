package mathutil

func GCD(r0, r1 int64) int64 {
	r0, r1 = AbsInt64(r0), AbsInt64(r1)
	if r1 > r0 {
		r0, r1 = r1, r0
	}

	for r1 != 0 {
		r0, r1 = r1, r0%r1
	}

	return r0
}

func LCM(a, b int64) int64 {
	return (a / GCD(a, b)) * b
}

func EGCD(r0, r1 int64) (int64, int64, int64) {
	var negA, negB bool = r0 < 0, r1 < 0
	var t0, s0, t1, s1 int64 = 0, 1, 1, 0
	r0, r1 = AbsInt64(r0), AbsInt64(r1)
	for r0 != 0 {
		quotient := r1 / r0
		r0, r1 = r1-quotient*r0, r0
		t0, t1 = t1, t0-t1*quotient
		s0, s1 = s1, s0-s1*quotient
	}
	if negA {
		t0 = -t0
	}
	if negB {
		s0 = -s0
	}
	return r1, t0, s0
}

func ModularInverse(x, n int64) int64 {
	g, s, _ := EGCD(x, n)
	if g == 1 {
		if s < 0 {
			return s + n
		} else {
			return s
		}
	}
	panic("Modular inverse does not exist")
}
