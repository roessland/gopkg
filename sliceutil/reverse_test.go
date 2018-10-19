package sliceutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInt64(t *testing.T) {
	s := []int64{1, 2, 3, 4}
	expected := []int64{4, 3, 2, 1}
	ReverseInt64(s)
	assert.True(t, EqualInt64(expected, s))
}
