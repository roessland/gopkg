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
