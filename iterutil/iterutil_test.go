package iterutil

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCartesianPower(t *testing.T) {
    var tuples <-chan []int64

    // (Z_1)**1 = { (0) }
    tuples = CartesianPower(1, 1)
    assert.Equal(t, []int64{0}, <-tuples, "they should be equal")
    assert.Nil(t, <-tuples, "channel should be closed")

    // (Z_1)**2 = { (0,0) }
    tuples = CartesianPower(1, 2)
    assert.Equal(t, []int64{0, 0}, <-tuples, "they should be equal")
    assert.Nil(t, <-tuples, "channel should be closed")

    // (Z_2)**2 = { (0,0), (0,1), (1,0), (1,1) }
    tuples = CartesianPower(2, 2)
    assert.Equal(t, []int64{0, 0}, <-tuples, "they should be equal")
    assert.Equal(t, []int64{0, 1}, <-tuples, "they should be equal")
    assert.Equal(t, []int64{1, 0}, <-tuples, "they should be equal")
    assert.Equal(t, []int64{1, 1}, <-tuples, "they should be equal")
    assert.Nil(t, <-tuples, "channel should be closed")

    // (Z_5)**1 = { (0), (1), (2), (3), (4) }
    tuples = CartesianPower(5, 1)
    assert.Equal(t, []int64{0}, <-tuples, "they should be equal")
    assert.Equal(t, []int64{1}, <-tuples, "they should be equal")
    assert.Equal(t, []int64{2}, <-tuples, "they should be equal")
    assert.Equal(t, []int64{3}, <-tuples, "they should be equal")
    assert.Equal(t, []int64{4}, <-tuples, "they should be equal")
    assert.Nil(t, <-tuples, "channel should be closed")
}
