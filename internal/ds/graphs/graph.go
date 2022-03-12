package graphs

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		fmt.Println("error: vertex already exists")
		return
	}

	g.vertices = append(g.vertices, &Vertex{key: k})
}

func (g *Graph) AddEdge(vertex, edge int) {
	v := g.GetVertex(vertex)
	e := g.GetVertex(edge)

	if v == nil || e == nil {
		fmt.Printf("Edge or vertex not found %d %d \n", edge, vertex)
		return
	}

	if contains(v.adjacent, edge) {
		fmt.Printf("Vertex %d already has edge %d \n", vertex, edge)
		return
	}

	v.adjacent = append(v.adjacent, e)
}

func (g *Graph) GetVertex(k int) *Vertex {
	for _, v := range g.vertices {
		if k == v.key {
			return v
		}
	}

	return nil
}

func (g *Graph) Print() {
	for _, e := range g.vertices {
		fmt.Printf("vertex %d -> edges: ", e.key)
		for _, v := range e.adjacent {
			fmt.Printf("%d ", v.key)
		}

		fmt.Println()
	}
}

func contains(vertices []*Vertex, k int) bool {
	for _, v := range vertices {
		if k == v.key {
			return true
		}
	}
	return false
}
