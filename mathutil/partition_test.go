package mathutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPartitionsInt(t *testing.T) {
	expectPartitionsMs := []MultisetIntInt{
		{4: 1},
		{3: 1, 1: 1},
		{2: 2},
		{2: 1, 1: 2},
		{1: 4},
	}
	actualPartitions := PartitionsInt(4)
	fmt.Println(actualPartitions)
	require.NotNil(t, actualPartitions)
	actualPartitionsMs := PartitionsIntToMultisetsIntInt(actualPartitions)
	for _, partition := range actualPartitionsMs {
		fmt.Println(partition)
	}
	require.Equal(t, len(expectPartitionsMs), len(actualPartitionsMs))
	for i, _ := range expectPartitionsMs {
		require.NotNil(t, actualPartitions[i])
		assert.True(t, actualPartitionsMs[i].Equal(expectPartitionsMs[i]))
	}

	assert.Equal(t, 11, len(PartitionsInt(6)))
	assert.Equal(t, 15, len(PartitionsInt(7)))
	assert.Equal(t, 22, len(PartitionsInt(8)))
	assert.Equal(t, 30, len(PartitionsInt(9)))
}

func TestNumPartitionsInt(t *testing.T) {
	expectPartitionsMs := []MultisetIntInt{
		{4: 1},
		{3: 1, 1: 1},
		{2: 2},
		{2: 1, 1: 2},
		{1: 4},
	}
	actualPartitions := PartitionsInt(4)
	fmt.Println(actualPartitions)
	require.NotNil(t, actualPartitions)
	actualPartitionsMs := PartitionsIntToMultisetsIntInt(actualPartitions)
	for _, partition := range actualPartitionsMs {
		fmt.Println(partition)
	}
	require.Equal(t, len(expectPartitionsMs), len(actualPartitionsMs))
	for i, _ := range expectPartitionsMs {
		require.NotNil(t, actualPartitions[i])
		assert.True(t, actualPartitionsMs[i].Equal(expectPartitionsMs[i]))
	}

	assert.Equal(t, 11, len(PartitionsInt(6)))
	assert.Equal(t, 15, len(PartitionsInt(7)))
	assert.Equal(t, 22, len(PartitionsInt(8)))
	assert.Equal(t, 30, len(PartitionsInt(9)))
}
