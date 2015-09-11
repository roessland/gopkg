package mathutil

func Factorial(n int64) int64 {
	if n > 20 {
		panic("Integer overflow in Factorial")
	}
	prod := int64(1)
	for i := int64(2); i <= n; i++ {
		prod *= i
	}
	return prod
}

func FactorialMod(n, p int64) int64 {
	prod := int64(1)
	for i := int64(2); i <= n; i++ {
		prod = (prod * i) % p
	}
	return prod
}

// Returns the number of zeroes at the end of the factorial n!.
// For example, 15! = 1307674368000, so it n! is divisible by 10^3,
// so the result must be 3.
func FactorialTrailingZeroes(n int) int {
	//               Num
	// 1344 / 5    = 268
	// 1344 / 5^2  =  53
	// 1344 / 5^3  =  10
	// 1344 / 5^4  =   2
	// Total       = 333
	numTens := 0
	for k := n / 5; k != 0; k /= 5 {
		numTens += k
	}
	return numTens
}
