package mathutil

import "testing"
import "github.com/stretchr/testify/assert"

func TestToDigits(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 3}, ToDigits(123, 10), "they should be equal")
	assert.Equal(t, []int64{1, 1, 1, 2, 0}, ToDigits(123, 3), "they should be equal")
	assert.Panics(t, func() {
		ToDigits(-123, 10)
	}, "negative number should panic")
	assert.Panics(t, func() {
		ToDigits(123, 1)
	}, "base 1 should panic")

}

func TestFromDigits(t *testing.T) {
	assert.Equal(t, int64(123), FromDigits([]int64{1, 2, 3}, 10), "they should be equal")
	assert.Equal(t, int64(123), FromDigits([]int64{1, 1, 1, 2, 0}, 3), "they should be equal")
	assert.Panics(t, func() {
		FromDigits([]int64{1, 1, 1, 2, 3}, 3)
	}, "digit not possible should panic")
}

func TestIsPandigital(t *testing.T) {
	assert.Equal(t, true, IsPandigital(123456789, 9), "they should be equal")
	assert.Equal(t, true, IsPandigital(12345678, 8), "they should be equal")
	assert.Equal(t, true, IsPandigital(1, 1), "they should be equal")
	assert.Equal(t, false, IsPandigital(12345678, 9), "they should be equal")
	assert.Equal(t, false, IsPandigital(123458789, 9), "they should be equal")
	assert.Equal(t, false, IsPandigital(123458789, 7), "they should be equal")
	assert.Equal(t, false, IsPandigital(510152025, 9), "they should be equal")
}

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome(1))
	assert.False(t, IsPalindrome(12))
	assert.True(t, IsPalindrome(121))
	assert.True(t, IsPalindrome(12355321))
}

func TestReverse(t *testing.T) {
	var base int64 = 10
	cases := [][]int64{
		{1, 1},
		{0, 0},
		{12, 21},
		{123, 321},
		{120, 21},
		{916413, 314619},
	}
	for _, c := range cases {
		assert.Equal(t, c[1], Reverse(c[0], base))
	}
}
