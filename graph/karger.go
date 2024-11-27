package graph

type (
	NodeID int
	EdgeID int
)

type graph struct {
	Vertices map[NodeID]map[EdgeID]struct{}
	Edges    map[EdgeID]*Edge
}

func newGraph(G Graph) *graph {
	g := &graph{
		Vertices: make(map[NodeID]map[EdgeID]struct{}),
		Edges:    make(map[EdgeID]*Edge),
	}
	for v, edges := range G.Vertices {
		g.Vertices[v] = make(map[EdgeID]struct{})
	}
}

type Graph struct {
	Vertices []NodeID
	Edges    []Edge
}

func (g *Graph) Clone() *Graph {
	gClone := &Graph{
		Vertices: make(map[NodeID]map[EdgeID]struct{}),
		Edges:    make(map[EdgeID]*Edge),
	}
	for v, edges := range g.Vertices {
		gClone.Vertices[v] = make(map[EdgeID]struct{})
		for e := range edges {
			gClone.Vertices[v][e] = struct{}{}
		}
	}
	for e, edge := range g.Edges {
		gClone.Edges[e] = edge
	}
	return gClone
}

type Edge struct {
	From   NodeID
	To     NodeID
	Weight float64
}

func Karger(graph Graph) []Edge {
	numVertices := len(graph.Vertices)
	// numEdges := len(graph.Edges)

	return nil
}

func contract(graph Graph) []Edge {
	numVertices := len(graph.Vertices)
	// numEdges := len(graph.Edges)
	return nil
}

