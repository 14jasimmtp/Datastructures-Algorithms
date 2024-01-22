package main

import "fmt"

type graph struct {
	vertices []any
	edges    map[any][]any
}

func (g *graph) addVertex(value any) {
	if g.contains(value) {
		return
	}

	g.vertices = append(g.vertices, value)
	g.edges[value] = []any{}
}

func (g *graph) addEdge(vertex1 any, vertex2 any) {
	if g.contains(vertex1) && g.contains(vertex2) {
		g.edges[vertex1] = append(g.edges[vertex1], vertex2)
		g.edges[vertex2] = append(g.edges[vertex2], vertex1)
	}
}

func (g *graph) print() {
	for key, value := range g.edges {
		fmt.Printf("%v : %v", key, value)
		fmt.Println()
	}
}

func (g *graph) contains(value any) bool {
	for _, val := range g.vertices {
		if val == value {
			return true
		}
	}
	return false
}

func (g *graph) BFS(startVertex any) {

	queue := []any{}
	visited := make(map[any]bool)
	queue = append(queue, startVertex)
	visited[startVertex] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		fmt.Print(" ", vertex)

		for _, neighbours := range g.edges[vertex] {
			if !visited[neighbours] {
				visited[neighbours] = true
				queue = append(queue, neighbours)
			}
		}
	}

}

func main() {
	gra := graph{
		edges: make(map[any][]any),
	}

	gra.addVertex(5)
	gra.addVertex(8)
	gra.addVertex(3)
	gra.addVertex("hg")
	gra.addVertex(4.43)
	gra.addEdge(5, 8)
	gra.addEdge("hg", 4.43)
	gra.addEdge(8, 4.43)
	gra.addEdge(3, "hg")
	gra.BFS(3)
	gra.print()
}
