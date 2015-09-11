package mathutil

// Calculates a**b for positive a and b
func Pow(a, b int64) int64 {
	p := int64(1)
	for b > 0 {
		if b&1 != 0 {
			p = p * a
		}
		b >>= 1
		a = a * a
		if a < 0 {
			panic("Integer overflow!")
		}
	}
	return p
}

// Calculates a**b mod m
func PowMod(a, b, m int64) int64 {
	a = a % m
	p := int64(1) % m
	for b > 0 {
		if b&1 != 0 {
			p = (p * a) % m
		}
		b >>= 1
		a = (a * a) % m
		if a < 0 {
			panic("Integer overflow!")
		}
	}
	return p
}
