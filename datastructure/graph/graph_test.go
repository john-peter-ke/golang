package graph

import "testing"

func TestGraph(t *testing.T) {
	t.Run("Test graph", func(t *testing.T) {
		v1 := 1
		g := graph{}
		g.AddVertex(v1)
		g.AddVertex(v1)
	})
}
