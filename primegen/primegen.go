package primegen

import (
	"math"
)

// PrimeMap returns an isPrime boolean map containing the numbers up to and
// including N. PrimeMap uses the sieve of Eratosthenes and has a running time
// of O(N).
func Map(N int) []bool {
	if N < 0 {
		return []bool{}
	}

	// Map starts out true
	A := make([]bool, N+1)
	for i := range A {
		A[i] = true
	}

	// The first two values aren't handles by the sieve, so we manually
	// specify that they aren't prime.
	if N >= 0 {
		A[0] = false
	}
	if N >= 1 {
		A[1] = false
	}

	// Sieve
	for i := 2; i <= int(math.Sqrt(float64(N))); i++ {
		if A[i] {
			for j := i * i; j <= N; j += i {
				A[j] = false
			}
		}
	}
	return A

}

// PrimesUpToSlice takes an isPrime boolean map, and returns a slice of the
// primes.
func Slice(isPrime []bool) []int {
	primes := make([]int, 0)
	for num, numIsPrime := range isPrime {
		if numIsPrime {
			primes = append(primes, num)
		}
	}
	return primes
}

// PrimesChan takes an isPrime boolean map, and sends the primes to a
// channel.
func Chan(isPrime []bool, primes chan<- int) {
	for num, numIsPrime := range isPrime {
		if numIsPrime {
			primes <- num
		}
	}
}
