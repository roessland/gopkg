package mathutil

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestPowMod(t *testing.T) {
	assert.Equal(t, int64(3403), PowMod(353, 434, 7654), "")
}

func TestPowModStrangeCase(t *testing.T) {
	assert.Equal(t, int64(3037001881), Pow[int64](55109, 2), "")
}

func TestPowModStrangeCases(t *testing.T) {
	for i := int64(1); i < 5520550; i += 100 {
		assert.Equal(t, i*i, Pow[int64](i, 2), "asdf")
	}
}
