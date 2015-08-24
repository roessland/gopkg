package mathutil

import "fmt"

// Choose calculate n choose k
func Choose(n, k int64) int64 {
	if n < k {
		return 0
	}
	if n-k < k {
		return Choose(n, n-k)
	}
	L := make([]int64, k+1)
	L[k] = 1
	for i := int64(0); i <= n-k; i++ {
		for j := k - 1; j >= 0; j-- {
			L[j] = L[j] + L[j+1]

			// Safeguard for integer overflow. Does not catch everything.
			if L[j] < 0 {
				panic(fmt.Sprintf("Integer overflow in Choose(%v, %v)", n, k))
			}
		}
	}
	return L[0]
}

// ChooseMod calculates n choose k mod m
func ChooseMod(n, k, m int64) int64 {
	if n < k {
		return 0
	}
	if n-k < k {
		return ChooseMod(n, n-k, m)
	}
	L := make([]int64, k+1)
	L[k] = 1
	for i := int64(0); i <= n-k; i++ {
		for j := k - 1; j >= 0; j-- {
			L[j] = L[j] + L[j+1]

			// Safeguard for integer overflow. Does not catch everything.
			if L[j] < 0 {
				panic(fmt.Sprintf("Integer overflow in ChooseMod(%v, %v, %v)", n, k, m))
				//c:panic("Integer overflow in Choose(n, k).")
			} else {
				L[j] %= m
			}
		}
	}
	return L[0]
}

// ChooseModPrime calculates n choose k mod p, for prime p
// See Lucas' theorem -- https://en.wikipedia.org/wiki/Lucas%27_theorem
func ChooseModPrime(n, k, p int64) int64 {
	if n < k {
		return 0
	}
	// Convert to base p
	ni := ToDigits(n, p)
	ki := ToDigits(k, p)

	prod := int64(1)
	for i := len(ki) - 1; i >= 0; i-- {
		prod = (prod * ChooseMod(ni[i], ki[i], p)) % p
	}
	return prod
}
