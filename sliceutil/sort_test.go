package sliceutil

import "testing"
import "github.com/stretchr/testify/assert"

func TestSortInt64(t *testing.T) {
	var slice []int64
	slice = []int64{3, 2, 1}
	SortInt64(slice)
	assert.Equal(t, []int64{1, 2, 3}, slice, "")
}
