package sliceutil

import "testing"
import "github.com/stretchr/testify/assert"

func TestIntersect(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 5}, Intersect([]int64{1, 2, 3, 4, 5}, []int64{1, 2, 5, 6, 7}), "")
}

func TestUnion(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7}, Union([]int64{1, 2, 3, 4, 5}, []int64{1, 2, 5, 6, 7}), "")
}
