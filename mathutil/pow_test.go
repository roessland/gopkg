package mathutil

import "testing"
import "github.com/stretchr/testify/assert"

func TestPowMod(t *testing.T) {
	assert.Equal(t, int64(3403), PowMod(353, 434, 7654), "")
}
