package mathutil

import "testing"
import "github.com/stretchr/testify/assert"

func TestChoose(t *testing.T) {
	assert.Equal(t, int64(120), Choose(10, 3), "simple binomial coefficient")
	assert.Equal(t, int64(10), Choose(5, 3), "simple binomial coefficient")
	assert.Equal(t, int64(286), Choose(13, 3), "simple binomial coefficient")
	assert.Equal(t, int64(286), Choose(13, 13-3), "simple binomial coefficient")
	assert.Panics(t, func() { Choose(1000, 50) }, "integer overflow")
	assert.Equal(t, int64(0), Choose(3, 10), "ways to choose 10 items out of 3")
}

func TestChooseMod(t *testing.T) {
	assert.Equal(t, int64(120), ChooseMod(10, 3, 155), "simple binomial coefficient")
	assert.Equal(t, int64(10), ChooseMod(5, 3, 11), "simple binomial coefficient")
	assert.Equal(t, int64(48675), ChooseMod(1000, 50, 454355), "large numbers")
	assert.Equal(t, int64(869659866), ChooseMod(10000, 764, 4543557657), "large mod")
	assert.Equal(t, int64(4026078894202900055), ChooseMod(10000, 764, 4543557657454355765), "very large mod A")
	assert.Equal(t, int64(4026078894202900055), ChooseMod(10000, 10000-764, 4543557657454355765), "very large mod symmetry")
	assert.Equal(t, int64(0), ChooseMod(3, 10, 5), "ways to choose 10 items out of 3")
}

func TestChooseModPrime(t *testing.T) {
	assert.Equal(t, int64(4), ChooseModPrime(588, 277, 5), "small prime")
	assert.Equal(t, int64(25), ChooseModPrime(6217032, 2484647, 107), "larger prime")
	assert.Equal(t, int64(0), ChooseModPrime(3, 10, 5), "ways to choose 10 items out of 3")
}
