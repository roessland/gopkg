package mathutil

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
