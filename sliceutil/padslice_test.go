package sliceutil

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPadLeft(t *testing.T) {
    assert.Equal(t, []int64{0,0,0,34,50,9}, ZeroPadLeftInt64([]int64{34,50,9}, 6), "they should be equal")
    assert.Equal(t, []int64{0,34,50,9}, ZeroPadLeftInt64([]int64{0,34,50,9}, 4), "they should be equal")
    assert.Equal(t, []int64{1,0,34,50,9}, ZeroPadLeftInt64([]int64{1,0,34,50,9}, 3), "they should be equal")
}
