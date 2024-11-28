package karger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
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

	rand.Seed(uint64(time.Now().UnixNano()))

	// Karger contraction is a probabilistic algorithm
	numFailures := 0
	numSuccesses := 0
	for i := 0; i < 100; i++ {
		result := Karger(g)
		if len(result.Edges) == 3 {
			numSuccesses++
		} else {
			numFailures++
		}
	}
	assert.Greater(t, numSuccesses, 10)
}
