package graph

import "testing"

func TestGraph(t *testing.T) {
	t.Run("Test graph", func(t *testing.T) {
		g := graph{}

		for i := 0; i < 5; i++ {
			g.AddVertex(i)
		}

		g.AddEdge(1, 2)
		g.AddEdge(1, 2)
		g.AddEdge(6, 2)
		g.AddEdge(3, 2)

		g.Print()
	})
}
