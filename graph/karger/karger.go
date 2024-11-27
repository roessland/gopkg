package karger

import (
	"github.com/roessland/gopkg/disjointset"
	"golang.org/x/exp/rand"
)

type Graph struct {
	Edges       []Edge
	NumVertices int
}

type Edge struct {
	A int
	B int
}

func NewUnweightedGraph(numVertices int) *Graph {
	return &Graph{
		Edges:       []Edge{},
		NumVertices: numVertices,
	}
}

func (g *Graph) AddEdge(a, b int) {
	g.Edges = append(g.Edges, Edge{A: a, B: b})
}

func (g *Graph) AddVertex() {
	g.NumVertices++
}

func Karger(g *Graph) Result {
	V := g.NumVertices
	E := len(g.Edges)
	edges := g.Edges
	ds := disjointset.Make(V)

	for V > 2 {
		// Pick random edge to contract
		ie := rand.Intn(E)
		from, to := edges[ie].A, edges[ie].B
		componentFrom, componentTo := ds.Find(from), ds.Find(to)

		// No point in contracting if the vertices are already in the same component
		if componentFrom == componentTo {
			continue
		}

		// Contract
		ds.Union(from, to)
		V--
	}

	// Find two remaining components
	var componentA, componentB int
	componentA = ds.Find(0)
	for i := 1; i < g.NumVertices; i++ {
		componentB = ds.Find(i)
		if componentB != componentA {
			break
		}
	}

	// Find sizes of the two components
	sizeA, sizeB := 0, 0
	for i := 0; i < g.NumVertices; i++ {
		if ds.Find(i) == componentA {
			sizeA++
		} else if ds.Find(i) == componentB {
			sizeB++
		} else {
			panic("Vertex not in any component")
		}
	}

	// Find edges between them
	cutEdges := []Edge{}
	for _, e := range g.Edges {
		componentA := ds.Find(e.A)
		componentB := ds.Find(e.B)
		if componentA != componentB {
			cutEdges = append(cutEdges, e)
		}
	}

	return Result{
		Edges: cutEdges,
		SizeA: sizeA,
		SizeB: sizeB,
	}
}

type Result struct {
	Edges        []Edge
	SizeA, SizeB int
}
