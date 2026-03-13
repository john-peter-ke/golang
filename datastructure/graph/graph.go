package graph

import "fmt"

type graph struct {
	vertices []*vertex
}

type vertex struct {
	key      int
	adjacent []*vertex
}

func (g *graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		fmt.Println("vertex key already exist")
	} else {
		g.vertices = append(g.vertices, &vertex{key: k})
	}
}

func (g *graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Println("invalid edges")
	} else if contains(fromVertex.adjacent, to) {
		fmt.Println("edge key already exist")
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		toVertex.adjacent = append(toVertex.adjacent, fromVertex)
	}
}

func (g *graph) getVertex(k int) *vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}

	return nil
}

func contains(s []*vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}

	return false
}

func (g *graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}

	fmt.Println()
}
