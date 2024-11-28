package karger

import (
	"fmt"
	"math/rand"

	"github.com/roessland/gopkg/disjointset"
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
	return contract(g, disjointset.Make(g.NumVertices), 2)
}

func contract(g *Graph, ds *disjointset.DisjointSet, t int) Result {
	if t < 2 {
		panic("t: Number of vertices must be at least 2")
	}
	V := ds.Count
	if V < 2 {
		panic("Number of vertices must be at least 2")
	}
	E := len(g.Edges)
	edges := g.Edges

	edgesContracted := 0
	for V > t {
		if edgesContracted >= E {
			panic("no more edges")
		}
		// Pick random edge to contract
		ie := rand.Intn(E)
		from, to := edges[ie].A, edges[ie].B

		// No point in contracting if the vertices are already in the same component
		if ds.Connected(from, to) {
			continue
		}

		// Contract
		edgesContracted++
		ds.Union(from, to)
		V--
	}

	// Find two remaining components
	componentA, componentB := ds.Find(0), ds.Find(0)
	for i := 0; i < g.NumVertices; i++ {
		if !ds.Connected(i, componentA) {
			componentB = ds.Find(i)
			break
		}
	}
	if componentA == componentB {
		panic("componentA == componentB")
	}

	//	Find sizes of the two components
	sizeA, sizeB := 0, 0
	for i := 0; i < g.NumVertices; i++ {
		if ds.Connected(i, componentA) {
			sizeA++
		} else {
			sizeB++
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

	if ds.Count <= 1 {
		panic(fmt.Sprintf("ds.Count <= 1: %d", ds.Count))
	}
	return Result{
		Edges: cutEdges,
		SizeA: sizeA,
		SizeB: sizeB,
		DS:    ds,
	}
}

type Result struct {
	DS    *disjointset.DisjointSet
	Edges []Edge
	SizeA int
	SizeB int
}
