package mathutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicePartitions(t *testing.T) {
	ofZero := [][][]int{}
	assert.Equal(t, ofZero, SlicePartitions([]int{}))

	ofOne := [][][]int{{{1}}}
	assert.Equal(t, ofOne, SlicePartitions([]int{1}))

	ofTwo := [][][]int{
		{{5, 4}},
		{{5}, {4}},
	}
	assert.Equal(t, ofTwo, SlicePartitions([]int{5, 4}))

	ofFive := [][][]int{
		{{1, 2, 3, 4, 5}},
		{{1, 2, 3, 4}, {5}},
		{{1, 2, 3}, {4, 5}},
		{{1, 2, 3}, {4}, {5}},
		{{1, 2}, {3, 4, 5}},
		{{1, 2}, {3, 4}, {5}},
		{{1, 2}, {3}, {4, 5}},
		{{1, 2}, {3}, {4}, {5}},
		{{1}, {2, 3, 4, 5}},
		{{1}, {2, 3, 4}, {5}},
		{{1}, {2, 3}, {4, 5}},
		{{1}, {2, 3}, {4}, {5}},
		{{1}, {2}, {3, 4, 5}},
		{{1}, {2}, {3, 4}, {5}},
		{{1}, {2}, {3}, {4, 5}},
		{{1}, {2}, {3}, {4}, {5}},
	}
	assert.Equal(t, 2*2*2*2, len(ofFive))
	assert.Equal(t, ofFive, SlicePartitions([]int{1, 2, 3, 4, 5}))

}
