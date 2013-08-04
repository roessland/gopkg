// Package sliceutil provides utility functions for slices
package sliceutil

// IntsEqual compares two slices of ints, and returns true only if all elements
// are equal.
func IntsEqual(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

// BoolsEqual compares two slices of bools, and returns true only if all elements
// are equal.
func BoolsEqual(a, b []bool) bool {
    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
