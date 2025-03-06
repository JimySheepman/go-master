package algorithm

import "fmt"

type graph func()

var graphAlgorithms = []graph{
	DepthFirstSearch,
	BreadthFirstSearch,
}

var graphAlgorithmsName = map[int]string{
	0: "Depth First Search",
	1: "Breadth First Search",
}

type Graph struct {
	vertices map[int][]int
}

func NewGraph() *Graph {
	return &Graph{vertices: make(map[int][]int)}
}

func (g *Graph) AddEdge(v, w int) {
	g.vertices[v] = append(g.vertices[v], w)
	g.vertices[w] = append(g.vertices[w], v) // This is for undirected graph
}

func (g *Graph) DFS(start int, visited map[int]bool) {
	visited[start] = true
	fmt.Printf("%d ", start)
	for _, v := range g.vertices[start] {
		if !visited[v] {
			g.DFS(v, visited)
		}
	}
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", s)

		for _, v := range g.vertices[s] {
			if !visited[v] {
				queue = append(queue, v)
				visited[v] = true
			}
		}
	}
}

func DepthFirstSearch() {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)
	visited := make(map[int]bool)
	fmt.Println("Depth-First Search starting from vertex 2:")
	g.DFS(2, visited)
}

func BreadthFirstSearch() {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	fmt.Println("Breadth-First Search starting from vertex 2:")
	g.BFS(2)
}

func PrintGraphAlgorithm() {
	for i, graphFunc := range graphAlgorithms {
		fmt.Println("Algorithm name:", graphAlgorithmsName[i])
		graphFunc()
	}
}
