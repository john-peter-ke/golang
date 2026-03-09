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
		fmt.Println("vertex already exist")
	}

	g.vertices = append(g.vertices, &vertex{key: k})

}

func (g *graph) AddEdge(from, to int) {

}

func contains(s []*vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}

	return false
}
