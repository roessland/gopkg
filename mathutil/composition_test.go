package mathutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompositionsOfZero(t *testing.T) {
	var expect [][]int
	actual := Compositions(0)
	require.Equal(t, expect, actual)
}

func TestCompositionsOfOne(t *testing.T) {
	expect := [][]int{{1}}
	actual := Compositions(1)
	require.Equal(t, expect, actual)
}

func TestCompositionsOfTwo(t *testing.T) {
	expect := [][]int{{2}, {1, 1}}
	actual := Compositions(2)
	require.Equal(t, expect, actual)
}

func TestCompositionsOfFive(t *testing.T) {
	expect := [][]int{
		{5},
		{4, 1},
		{3, 2},
		{3, 1, 1},
		{2, 3},
		{2, 2, 1},
		{2, 1, 2},
		{2, 1, 1, 1},
		{1, 4},
		{1, 3, 1},
		{1, 2, 2},
		{1, 2, 1, 1},
		{1, 1, 3},
		{1, 1, 2, 1},
		{1, 1, 1, 2},
		{1, 1, 1, 1, 1},
	}
	actual := Compositions(5)
	require.Equal(t, expect, actual)
}

func TestCompositionsOfTen(t *testing.T) {
	expectedLen := 512
	actual := Compositions(10)
	require.Equal(t, expectedLen, len(actual))
}
