package digraph

import "fmt"
import "testing"
import "github.com/stretchr/testify/assert"

func TestPriorityQueue(t *testing.T) {
	assert.Equal(t, 1, 1)
	graph := Graph{[]Node{
		Node{[]Edge{Edge{To: 1, Weight: 5}, Edge{To: 2, Weight: 1}}},
		Node{[]Edge{Edge{To: 2, Weight: 1}, Edge{To: 3, Weight: 5}}},
		Node{[]Edge{Edge{To: 1, Weight: 1}, Edge{To: 3, Weight: 2}}},
		Node{[]Edge{}},
	}}
	_, prev := Dijkstra(graph, 0)
	fmt.Printf("%v\n", prev)
}
