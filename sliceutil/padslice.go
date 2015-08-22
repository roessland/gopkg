package  sliceutil

// Pad a slice to a certain length with zeros to the left.
// If the slice is longer than the desired length, nothing is done.
func ZeroPadLeftInt64(slice []int64, desiredLength int64) []int64 {
    currentLength := int64(len(slice))
    if currentLength >= desiredLength {
        return slice
    } else {
        return append(make([]int64, desiredLength - currentLength), slice...)
    }
}
