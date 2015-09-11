package mathutil

import "testing"
import "github.com/stretchr/testify/assert"

func TestFactorial(t *testing.T) {
	assert.Equal(t, int64(2432902008176640000), Factorial(20), "large number")
	assert.Panics(t, func() { Factorial(21) }, "overflowing number")
	assert.Equal(t, int64(1), Factorial(0), "zero factorial")
}

func TestFactorialMod(t *testing.T) {
	assert.Equal(t, int64(47), FactorialMod(30, 73), "")
	assert.Equal(t, int64(17823357210954), FactorialMod(30, 65465477433546), "")
}

func TestFactorialTrailingZeroes(t *testing.T) {
	assert.Equal(t, 0, FactorialTrailingZeroes(0), "")
	assert.Equal(t, 0, FactorialTrailingZeroes(1), "")
	assert.Equal(t, 0, FactorialTrailingZeroes(2), "")
	assert.Equal(t, 0, FactorialTrailingZeroes(3), "")
	assert.Equal(t, 0, FactorialTrailingZeroes(4), "")
	assert.Equal(t, 1, FactorialTrailingZeroes(5), "")
	assert.Equal(t, 1, FactorialTrailingZeroes(6), "")
	assert.Equal(t, 1, FactorialTrailingZeroes(7), "")
	assert.Equal(t, 1, FactorialTrailingZeroes(8), "")
	assert.Equal(t, 1, FactorialTrailingZeroes(9), "")
	assert.Equal(t, 2, FactorialTrailingZeroes(10), "")
	assert.Equal(t, 2, FactorialTrailingZeroes(14), "")
	assert.Equal(t, 3, FactorialTrailingZeroes(15), "")
	assert.Equal(t, 22, FactorialTrailingZeroes(99), "")
	assert.Equal(t, 333, FactorialTrailingZeroes(1344), "")
}
