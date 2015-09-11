package mathutil

import "testing"
import "github.com/stretchr/testify/assert"

func TestPartition(t *testing.T) {
	var want, found [][]int64
	want = [][]int64{[]int64{1, 5}, []int64{2, 4}}
	found = Partition([]int64{1, 2, 4, 5}, 6, 2, 0)
	assert.Equal(t, want, found, "")

	want = [][]int64{[]int64{2, 18}, []int64{7, 13}}
	found = Partition([]int64{2, 7, 13, 18, 23, 25}, 20, 2, 0)
	assert.Equal(t, want, found, "")

	want = [][]int64{[]int64{3, 7, 109, 673}, []int64{5, 7, 107, 673}}
	found = Partition([]int64{2, 3, 5, 7, 11, 13, 17, 23, 29, 31, 37, 107, 109, 673}, 792, 4, 0)
	assert.Equal(t, want, found, "")
}
