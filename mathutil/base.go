package mathutil

import "github.com/roessland/gopkg/sliceutil"

func ToDigits(n, base int64) []int64 {
	if n < 0 {
		panic("negative number was provided")
	}
	if base <= 1 {
		panic("base must be 2 or higher")
	}

	digits := []int64{}

	for n > 0 {
		digits = append([]int64{n % base}, digits...)
		n = n / base
	}
	return digits
}

func FromDigits(digits []int64, base int64) int64 {
	var n int64
	for i, b := len(digits)-1, int64(1); i >= 0; i, b = i-1, b*base {
		if digits[i] >= base {
			panic("digit does not exist in this base")
		}
		n += digits[i] * b
	}
	return n
}

func IsPandigital(n int64, length int64) bool {
	digits := ToDigits(n, 10)
	if int64(len(digits)) != length {
		return false
	}
	counts := make([]int32, length)
	for _, digit := range digits {
		if digit == 0 || digit > length {
			return false
		}
		counts[digit-1]++
	}
	for i := int64(0); i < length; i++ {
		if counts[i] != 1 {
			return false
		}
	}
	return true
}

func IsPalindrome(n int64) bool {
	digits := ToDigits(n, 10)
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}

func Reverse(n int64, base int64) int64 {
	digits := ToDigits(n, 10)
	sliceutil.ReverseInt64(digits)
	return FromDigits(digits, 10)
}
