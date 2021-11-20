package mathutil

// MulMod calculates a*b mod m
func MulMod(a, b, m int64) int64 {
	var sum int64 = 0
	a, b = a%m, b%m
	for b != 0 {
		if b&1 == 1 {
			sum = (sum + a) % m
		}
		a, b = (2*a)%m, b>>1
	}
	return sum
}
