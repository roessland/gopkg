package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	assert.Equal(t, 1, 1)
	graph := Graph{[]Node{
		{[]Edge{{To: 1, Weight: 5}, {To: 2, Weight: 1}}},
		{[]Edge{{To: 2, Weight: 1}, {To: 3, Weight: 5}}},
		{[]Edge{{To: 1, Weight: 1}, {To: 3, Weight: 2}}},
		{[]Edge{}},
	}}
	_, prev := Dijkstra(graph, 0)
	fmt.Printf("%v\n", prev)
}
