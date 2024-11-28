package disjointset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	ds := Make(10)
	assert.Equal(t, 9, ds.Find(9))
}

func TestClone(t *testing.T) {
	ds := Make(10)
	assert.Equal(t, 9, ds.Find(9))

	dsClone := ds.Clone()
	assert.Equal(t, 9, dsClone.Find(9))
}

func TestUnionFind(t *testing.T) {
	ds := Make(10) // (8 9 1 2) (0 3 5 4) 7 5
	assert.Equal(t, 10, ds.Count)
	ds.Union(8, 9)
	assert.Equal(t, 9, ds.Count)
	ds.Union(1, 8)
	assert.Equal(t, 8, ds.Count)
	ds.Union(2, 1)
	assert.Equal(t, 7, ds.Count)
	ds.Union(0, 3)
	ds.Union(0, 4)
	ds.Union(0, 5)
	assert.Equal(t, 4, ds.Count)

	assert.True(t, ds.Connected(8, 9))
	assert.Equal(t, ds.Find(8), ds.Find(1))
	assert.Equal(t, ds.Find(8), ds.Find(2))
	assert.NotEqual(t, ds.Find(8), ds.Find(5))
	assert.Equal(t, ds.Find(3), ds.Find(5))

	fmt.Printf("%v\n", ds.rank)
}
