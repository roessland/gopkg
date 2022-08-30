package primegen

import (
	"math"
)

// Primes below 1000-ish
var Smallprimes []int64

// PrimeMap returns an isPrime boolean map containing the numbers up to and
// including N. PrimeMap uses the sieve of Eratosthenes and has a running time
// of O(N).
func Map(N int64) []bool {
	if N < 0 {
		return []bool{}
	}

	// Map starts out true
	A := make([]bool, N+1)
	for i := range A {
		A[i] = true
	}

	// The first two values aren't handled by the sieve, so we manually
	// specify that they aren't prime.
	if N >= 0 {
		A[0] = false
	}
	if N >= 1 {
		A[1] = false
	}

	// Sieve
	for i := int64(2); i <= int64(math.Sqrt(float64(N))); i++ {
		if A[i] {
			for j := i * i; j <= N; j += i {
				A[j] = false
			}
		}
	}
	return A

}

// FactorsMap generates prime factorization for numbers <= N
func FactorsMap(N int64) ([]bool, map[int64][]int64, map[int64][]int64) {
	if N < 0 {
		panic("makes no sense")
	}

	// Map starts out true
	A := make([]bool, N+1)
	for i := range A {
		A[i] = true
	}

	// Prime factors and multiplicity
	p := make(map[int64][]int64)
	k := make(map[int64][]int64)

	// The first two values aren't handled by the sieve, so we manually
	// specify that they aren't prime.
	if N >= 0 {
		A[0] = false
	}
	if N >= 1 {
		A[1] = false
	}

	// Sieve
	for i := int64(2); i <= N; i++ {
		if A[i] {
			// i is prime, with only one factor -- itself
			p[i] = append(p[i], i)
			k[i] = append(k[i], 1)

			// Add i as factor to all multiples of i
			for j := 2 * i; j <= N; j += i {
				A[j] = false

				// Find multiplicity
				multiplicity := int64(0)
				n := j
				for {
					if n%i == 0 {
						multiplicity++
					} else {
						break
					}
					n = n / i
				}
				p[j] = append(p[j], i)
				k[j] = append(k[j], multiplicity)
			}
		}
	}

	return A, p, k
}

type PrimeFactor struct {
	P int64
	A int
}

// FactorsSlice generates prime factorizations for numbers <= N,
// multiplicity included.
func FactorsSlice(N int64) ([]bool, [][]PrimeFactor) {
	if N < 0 {
		panic("makes no sense")
	}

	// Map starts out true
	isPrime := make([]bool, N+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	// Prime factors and multiplicity
	pfs := make([][]PrimeFactor, N+1)

	// The first two values aren't handled by the sieve, so we manually
	// specify that they aren't prime.
	if N >= 0 {
		isPrime[0] = false
	}
	if N >= 1 {
		isPrime[1] = false
	}

	// Sieve
	for i := int64(2); i <= N; i++ {
		if isPrime[i] {
			// i is prime, with only one factor -- itself
			pfs[i] = []PrimeFactor{{P: i, A: 1}}

			// Add i as factor to all multiples of i
			for j := 2 * i; j <= N; j += i {
				isPrime[j] = false

				// Find multiplicity
				multiplicity := 0
				n := j
				for {
					if n%i == 0 {
						multiplicity++
					} else {
						break
					}
					n = n / i
				}
				pfs[j] = append(pfs[j], PrimeFactor{P: i, A: multiplicity})
			}
		}
	}

	return isPrime, pfs
}

// The number of prime numbers less than or equal to n
func Pi(isPrime []bool, n int64) int64 {
	if n >= int64(len(isPrime)) {
		panic("isPrime boolean map was too small to decide")
	}
	count := int64(0)
	for num, numIsPrime := range isPrime {
		if int64(num) > n {
			break
		}

		if numIsPrime {
			count++
		}
	}
	return count
}

// PrimesUpToSlice takes an isPrime boolean map, and returns a slice of the
// primes.
func SliceFromMap(isPrime []bool) []int64 {
	primes := make([]int64, 0)
	for num, numIsPrime := range isPrime {
		if numIsPrime {
			primes = append(primes, int64(num))
		}
	}
	return primes
}

func init() {
	Smallprimes = SliceFromMap(Map(1000))
}
