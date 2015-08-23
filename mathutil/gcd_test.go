package mathutil

import "testing"
import "github.com/stretchr/testify/assert"

func TestGCD(t *testing.T) {
    assert.Equal(t, int64(3), GCD(453242343, 442143147), "they should be equal")
    assert.Equal(t, int64(3), GCD(3, 0), "they should be equal")
    assert.Equal(t, int64(3), GCD(0, 3), "they should be equal")
    assert.Equal(t, int64(3), GCD(3, 3), "they should be equal")
}
