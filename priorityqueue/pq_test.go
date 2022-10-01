package priorityqueue_test

import (
	"github.com/roessland/gopkg/priorityqueue"
	"github.com/stretchr/testify/require"
	"math/rand"
	"sort"
	"testing"
)

func TestPQ_1(t *testing.T) {
	pq := priorityqueue.New[string]()
	pq.Push("banana", 5)
	require.Equal(t, "banana", pq.Pop())
}

func TestPQ_2(t *testing.T) {
	pq := priorityqueue.New[string]()
	pq.Push("banana", 5)
	pq.Push("pear", 4)
	require.Equal(t, "pear", pq.Pop())
	require.Equal(t, "banana", pq.Pop())
}

func TestPQ_3(t *testing.T) {
	pq := priorityqueue.New[string]()
	pq.Push("pear", 4)
	pq.Push("banana", 5)
	require.Equal(t, "pear", pq.Pop())
	require.Equal(t, "banana", pq.Pop())
}

func TestPQ_4(t *testing.T) {
	pq := priorityqueue.New[int]()
	var expected []int
	for i := 0; i < 100000; i++ {
		val := rand.Intn(10000)
		pq.Push(val, float64(val))
		expected = append(expected, val)
	}

	var actual []int
	sort.Ints(expected)
	for pq.Len() > 0 {
		val := pq.Pop()
		actual = append(actual, val)
	}
	require.EqualValues(t, expected, actual)
}
