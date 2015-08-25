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
