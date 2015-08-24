package mathutil

import "testing"
import "github.com/stretchr/testify/assert"

func TestMaxInt64(t *testing.T) {
	assert.Equal(t, int64(0), MaxInt64(0, -100), "they should be equal")
	assert.Equal(t, int64(0), MaxInt64(0, 0), "they should be equal")
}

func TestAbsInt64(t *testing.T) {
	assert.Equal(t, int64(0), AbsInt64(0), "zero")
	assert.Equal(t, int64(2), AbsInt64(2), "positive number")
	assert.Equal(t, int64(2), AbsInt64(-2), "negative number")
}
