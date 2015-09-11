package digraph

import "math"
import "container/heap"

type Graph struct {
	Nodes []Node
}

type Node struct {
	Neighbors []Edge
}

type Edge struct {
	To     int
	Weight float64
}

type Item struct {
	value    int
	priority float64 // Distance from source to this node
	index    int     // Index of the item in the heap
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Finds the item with a certain value and changes its priority.
// Assumes unique values among all heap members.
func (pq *PriorityQueue) FindUpdate(value int, priority float64) {
	var i int
	for i = 0; i < len(*pq); i++ {
		if (*pq)[i].value == value {
			(*pq)[i].priority = priority
			heap.Fix(pq, (*pq)[i].index)
			return
		}
	}
	panic("Couldn't find and update node -- no node with that value")
}

func Dijkstra(graph Graph, source int) ([]float64, []int) {
	numNodes := len(graph.Nodes)
	//numEdges := len(graph.Edges)

	// Initialize distances to infinity
	dist := make([]float64, numNodes)
	for i, _ := range dist {
		dist[i] = math.Inf(1)
	}
	dist[source] = 0

	// Initialize pointers to previous nodes.
	// -1 means undefined.
	prev := make([]int, numNodes)
	for i, _ := range prev {
		prev[i] = -1
	}

	Q := make(PriorityQueue, numNodes)
	for i, _ := range graph.Nodes {
		Q[i] = &Item{
			value:    i,
			priority: dist[i],
			index:    i,
		}
	}
	heap.Init(&Q)

	for len(Q) > 0 {
		u := heap.Pop(&Q).(*Item).value
		for _, uv := range graph.Nodes[u].Neighbors {
			v := uv.To
			alt := dist[u] + uv.Weight
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				Q.FindUpdate(v, alt)
			}
		}
	}
	return dist, prev
}
