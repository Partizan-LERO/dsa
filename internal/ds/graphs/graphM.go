package graphs

import "fmt"

type GraphM struct {
	vertices map[int]*VertexM
}

type VertexM struct {
	adjacent map[int]*VertexM
}

func NewGraph() GraphM {
	fmt.Println("New Graph created!")
	return GraphM{
		vertices: make(map[int]*VertexM),
	}
}

func (g *GraphM) AddVertex(k int) {
	if _, ok := g.vertices[k]; ok {
		fmt.Println("Vertex already exists")
		return
	}

	g.vertices[k] = &VertexM{}
}

func (g *GraphM) AddEdge(vertex, edge int) {
	_, ok := g.vertices[vertex]
	if !ok {
		fmt.Printf("Vertex not found %d \n", vertex)
		return
	}

	e, ok := g.vertices[edge]
	if !ok {
		fmt.Printf("Edge not found %d \n", edge)
		return
	}

	if len(g.vertices[vertex].adjacent) == 0 {
		g.vertices[vertex].adjacent = make(map[int]*VertexM)
	}

	g.vertices[vertex].adjacent[edge] = e
}

func (g *GraphM) Print() {
	for vertex, _ := range g.vertices {
		fmt.Printf("vertex %d -> edges: ", vertex)
		for edge, _ := range g.vertices[vertex].adjacent {
			fmt.Printf("%d ", edge)
		}
		fmt.Println()
	}
}
