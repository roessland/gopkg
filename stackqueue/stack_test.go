package stackqueue_test

import (
	"github.com/roessland/gopkg/stackqueue"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStack_1(t *testing.T) {
	pq := stackqueue.New[string]()
	pq.Push("banana")
	require.Equal(t, "banana", pq.Pop())
}

func TestStack_2(t *testing.T) {
	pq := stackqueue.New[string]()
	pq.Push("banana")
	pq.Push("pear")
	require.Equal(t, "pear", pq.Pop())
	require.Equal(t, "banana", pq.Pop())
}

func TestStack_3(t *testing.T) {
	pq := stackqueue.New[string]()
	pq.Push("pear")
	pq.Push("banana")
	require.Equal(t, "banana", pq.Pop())
	require.Equal(t, "pear", pq.Pop())
}
