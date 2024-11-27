package karger

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	                  .─.
	                 ( 1 )
	                ┌─`─'─┐
	              ┌─┘ │ │ └─┐
	            ┌─┘  ┌┘ └┐  └─┐
	         ┌──┘    │   │    └─.─.
	   .─.┌──┘      ┌┴───┴┬────( 2 )
	  ( 0 )┬────────┤     └┐  ┌─`─'
	   `─' └───┐   ┌┘      └┬─┘  │
	    │└┐    └───┤      ┌─┴┐   │
	    │ │       ┌┴───┐┌─┘  └┐  └┐
	   ┌┘ └┐      │   ┌┴┴─┐   └┐  │
	   │   │     ┌┘ ┌─┘   └──┐ └┐ │
	   │   └┐   ┌┘┌─┘        └──┼.─.
	   │    │  ┌┼─┘        ┌────( 3 )
	   │    .─.┴┘  ┌───────┘     `─'
	  ┌┘   ( 4 )───┘              │
	  │     `─'                   │
	  │        ╲                 ┌┘
	  │         .─.              │
	 ┌┘        ( 5 )─────┐      .─.
	 │      ┌──┘`─'┬┬────┴─────( 6 )
	.─. ┌───┘   ┌┘ │└┐        ┌─`─'

( 9 )┘───────┼──┘ └┐     ┌─┘ ┌┘

	─  ┐──┐     │     └┐  ┌─┘   │
	   └┐ └──┐ ┌┘      └┬─┘     │
	    └┐   └─┼┐     ┌─┴┐     ┌┘
	     └┐   ┌┘└──┐┌─┘  └┐    │
	      │   │   ┌┴┴─┐   └┐  ┌┘
	      └┐ ┌┘ ┌─┘   └──┐ └┐ │
	       │ │┌─┘        └──.─.
	       .─.┘     ┌──────( 7 )
	      ( 8 )─────┘       `─'
	       `─'
*/
func TestKarger(t *testing.T) {
	g := NewUnweightedGraph(10)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 0)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(3, 0)
	g.AddEdge(3, 1)
	g.AddEdge(4, 1)
	g.AddEdge(4, 2)

	g.AddEdge(0+5, 1+5)
	g.AddEdge(1+5, 2+5)
	g.AddEdge(2+5, 3+5)
	g.AddEdge(3+5, 4+5)
	g.AddEdge(4+5, 0+5)
	g.AddEdge(0+5, 2+5)
	g.AddEdge(0+5, 3+5)
	g.AddEdge(1+5, 3+5)
	g.AddEdge(1+5, 4+5)
	g.AddEdge(3+5, 0+5)
	g.AddEdge(3+5, 1+5)
	g.AddEdge(4+5, 1+5)
	g.AddEdge(4+5, 2+5)

	g.AddEdge(0, 9)
	g.AddEdge(4, 5)
	g.AddEdge(3, 6)

	result := Karger(g)
	sort.Slice(result.Edges, func(i, j int) bool {
		return result.Edges[i].A < result.Edges[j].A
	})
	assert.Equal(t, 3, len(result.Edges))

	assert.Equal(t, 0, result.Edges[0].A)
	assert.Equal(t, 9, result.Edges[0].B)

	assert.Equal(t, 3, result.Edges[1].A)
	assert.Equal(t, 6, result.Edges[1].B)

	assert.Equal(t, 4, result.Edges[2].A)
	assert.Equal(t, 5, result.Edges[2].B)

	assert.Equal(t, 5, result.SizeA)
	assert.Equal(t, 5, result.SizeB)
}
