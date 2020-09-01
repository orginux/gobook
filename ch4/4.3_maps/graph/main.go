package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("one", "two")
	fmt.Println(graph)
	addEdge("two", "three")
	fmt.Println(graph)

	fmt.Println(hasEdge("one", "two"))
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
