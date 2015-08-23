package sliceutil

import "testing"

func TestEqualInt_different_length(t *testing.T) {
    a := []int{1, 2, 3}
    b := []int{1, 2}
    if EqualInt(a, b) {
        t.Errorf("IntsEqual(%v, %v) == true, want false")
    }
}

func TestEqualInt_same_length_but_different(t *testing.T) {
    a := []int{1, 2, 3}
    b := []int{1, 2, 4}
    if EqualInt(a, b) {
        t.Errorf("IntsEqual(%v, %v) == true, want false")
    }
}

func TestEqualInt_same(t *testing.T) {
    a := []int{1, 2, 3, 1}
    b := []int{1, 2, 3, 1}
    if !EqualInt(a, b) {
        t.Errorf("IntsEqual(%v, %v) == false, want true")
    }
}
