package iterutil

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCartesianPower(t *testing.T) {
    // (Z_1)**0 is an empty set
    tuples := CartesianPower(1, 0)
    emptySet <- tuples
    assert.Equal(t, emptySet, nil, "they should be equal")
}
