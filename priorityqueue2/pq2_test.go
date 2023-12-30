package priorityqueue2_test

import (
	"math"
	"math/rand"
	"sort"
	"testing"

	"github.com/roessland/gopkg/priorityqueue2"
	"github.com/stretchr/testify/require"
)

func TestPQ_1(t *testing.T) {
	pq := priorityqueue2.New[string, int]()
	pq.Push("banana", 5)
	require.Equal(t, "banana", pq.Pop())
}

func TestPQ_2(t *testing.T) {
	pq := priorityqueue2.New[string, int]()
	pq.Push("banana", 5)
	pq.Push("pear", 4)
	require.Equal(t, "pear", pq.Pop())
	require.Equal(t, "banana", pq.Pop())
}

func TestPQ_3(t *testing.T) {
	pq := priorityqueue2.New[string, int]()
	pq.Push("pear", 4)
	pq.Push("banana", 5)
	require.Equal(t, "pear", pq.Pop())
	require.Equal(t, "banana", pq.Pop())
}

func TestPQ_4(t *testing.T) {
	pq := priorityqueue2.New[int, float64]()
	var expected []int
	for i := 0; i < 10000; i++ {
		val := rand.Intn(5000)
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

// Using NaN really messes up the heap ordering since NaN < x is false for all x.
// Don't use NaN as a priority value.
func TestPQ_NaN_Screws_Up_Order(t *testing.T) {
	pq := priorityqueue2.New[string, float64]()
	pq.Push("one", 1.0)
	pq.Push("NaN", math.NaN())
	pq.Push("two", 2.0)
	pq.Push("+inf", math.Inf(1))
	pq.Push("-inf", math.Inf(-1))
	require.Equal(t, "one", pq.Pop())
	require.Equal(t, "-inf", pq.Pop())
	require.Equal(t, "+inf", pq.Pop())
	require.Equal(t, "two", pq.Pop())
	require.Equal(t, "NaN", pq.Pop())
}

// Using string priorityies might be useful for... something?
func TestPQ_FunTypes(t *testing.T) {
	pq := priorityqueue2.New[string, string]()
	pq.Push("banana_5", "5")
	pq.Push("banana_empty", "")
	pq.Push("banana_44", "4444444444")
	require.Equal(t, "banana_empty", pq.Pop())
	require.Equal(t, "banana_44", pq.Pop())
	require.Equal(t, "banana_5", pq.Pop())

	// "" < "5" < "4444444444"
}
