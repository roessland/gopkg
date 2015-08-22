package stringutil

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPadLeft(t *testing.T) {
    assert.Equal(t, "000日本語", PadLeft("日本語", 6, '0'), "they should be equal")
    assert.Equal(t, "まままま日本語", PadLeft("日本語", 7, 'ま'), "they should be equal")
    assert.Equal(t, "    ", PadLeft("", 4, ' '), "they should be equal")
}
