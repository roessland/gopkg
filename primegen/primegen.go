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
