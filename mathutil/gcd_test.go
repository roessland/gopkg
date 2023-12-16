package mathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	assert.Equal(t, int64(3), GCD(453242343, 442143147), "straight case")
	assert.Equal(t, int64(3), GCD(3, 0), "zero")
	assert.Equal(t, int64(3), GCD(0, 3), "zero")
	assert.Equal(t, int64(3), GCD(3, 3), "equal numbers")
	assert.Equal(t, int64(20), GCD(20, 100), "one is a multiplum of the other")
	assert.Equal(t, int64(20), GCD(-20, 100), "one is negative")
	assert.Equal(t, int64(20), GCD(20, -100), "one is negative")
	assert.Equal(t, int64(20), GCD(-20, -100), "both are negative")
}

func TestLCM(t *testing.T) {
	assert.Equal(t, int64(630), LCM(35, 90), "")
	assert.Equal(t, int64(0), LCM(1, 0), "")
}

func TestEGCD(test *testing.T) {
	var r, s, t, r_, s_, t_, a, b int64

	r, s, t = EGCD(3, 0)
	assert.True(test, r == int64(3) && s == int64(1) && t == int64(0), "zero value")

	r, s, t = EGCD(0, 3)
	assert.True(test, r == int64(3) && s == int64(0) && t == int64(1), "zero value")

	a, b = -2, -6
	r_, s_, t_ = 2, -1, 0
	r, s, t = EGCD(a, b)
	assert.True(test, r == r_ && s == s_ && t == t_, "negative values")

	a, b = -2, 5
	r_, s_, t_ = 1, 2, 1
	r, s, t = EGCD(a, b)
	assert.True(test, r == r_ && s == s_ && t == t_, "one negative value")

	a, b = 65, 40
	r_, s_, t_ = 5, -3, 5
	r, s, t = EGCD(a, b)
	assert.True(test, r == r_ && s == s_ && t == t_, "ordinary")

	a, b = 1239, 735
	r_, s_, t_ = 21, -16, 27
	r, s, t = EGCD(a, b)
	assert.True(test, r == r_ && s == s_ && t == t_, "ordinary")
}

func TestModularInverse(t *testing.T) {
	assert.Equal(t, int64(9), ModularInverse(5, 11), "prime modulo")
	assert.Equal(t, int64(5), ModularInverse(9, 11), "prime modulo")
	assert.Equal(t, int64(8), ModularInverse(7, 11), "prime modulo")
	assert.Panics(t, func() { ModularInverse(8, 12) }, "does not exist")
	assert.Equal(t, int64(5), ModularInverse(5, 12), "composite modulo")
	assert.Panics(t, func() { ModularInverse(0, 37) }, "0 has no inverse")
}
