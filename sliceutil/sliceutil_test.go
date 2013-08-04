package sliceutil

import "testing"

func TestIntsEqual_different_length(t *testing.T) {
    a := []int{1, 2, 3}
    b := []int{1, 2}
    if IntsEqual(a, b) {
        t.Errorf("IntsEqual(%v, %v) == true, want false")
    }
}

func TestIntsEqual_same_length_but_different(t *testing.T) {
    a := []int{1, 2, 3}
    b := []int{1, 2, 4}
    if IntsEqual(a, b) {
        t.Errorf("IntsEqual(%v, %v) == true, want false")
    }
}

func TestIntsEqual_same(t *testing.T) {
    a := []int{1, 2, 3, 1}
    b := []int{1, 2, 3, 1}
    if !IntsEqual(a, b) {
        t.Errorf("IntsEqual(%v, %v) == false, want true")
    }
}
