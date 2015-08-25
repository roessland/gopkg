package sliceutil

import "testing"
import "github.com/stretchr/testify/assert"

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

func TestStrictlyIncreasingInt64(t *testing.T) {
	assert.Equal(t, true, StrictlyIncreasingInt64([]int64{1, 2, 3, 4, 5}), "")
	assert.Equal(t, false, StrictlyIncreasingInt64([]int64{1, 2, 3, 4, 4}), "")
	assert.Equal(t, true, StrictlyIncreasingInt64([]int64{1}), "sequence length one")
	assert.Equal(t, true, StrictlyIncreasingInt64([]int64{}), "empty sequence")
}

func TestProductInt64(t *testing.T) {
	assert.Equal(t, int64(38500), ProductInt64([]int64{5, 14, 55, 10}), "")
}
