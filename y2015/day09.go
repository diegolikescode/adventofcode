package y2015

import (
	"fmt"
)

type graph struct {
	size          int
	edges         map[string]map[string]int
	// possiblePaths map[[]string]int
}

func (g *graph) printGraph() {
	for f, toMap := range g.edges {
		for to, distance := range toMap {
			fmt.Printf("from %s to %s distance = %d\n", f, to, distance)
		}
	}
	fmt.Printf("size of graph: %d | possible combinations: %d\n", g.size, Factorial(g.size))
}

func (g *graph) addVertice(v string) bool {
	if g.edges[v] == nil {
		g.edges[v] = make(map[string]int)
		g.size++
		return true
	}
	return false
}

func (g *graph) addEdge(from, to string, distance int) bool {
	if from == to {
		return false
	}

	g.addVertice(from)
	g.addVertice(to)
	if g.edges[from][to] == 0 || g.edges[to][from] == 0 {
		g.edges[from][to] = distance
		g.edges[to][from] = distance
		return true
	}
	return false
}

// func (g *graph) bruteForce() {
// 	visited := []string{}
// }

func (m SharedType) Day09_A(testFile string) {
	lines := ReturnFileByLine(testFile)
	g := &graph{
		size:  0,
		edges: make(map[string]map[string]int),
	}

	for _, l := range lines {
		var from, to string
		var distance int
		fmt.Sscanf(l, "%s to %s = %d", &from, &to, &distance)
		g.addEdge(from, to, distance)
	}
	// g.printGraph()
	// g.bruteForce()
}
