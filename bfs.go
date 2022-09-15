package main

import "fmt"

func main() {
	// undirected graph
	edges := [][]int{
		{0, 1},
		{0, 3},
		{1, 2},
		{1, 5},
		{2, 3},
		{2, 4},
		{2, 5},
		{3, 4},
		{4, 5},
	}

	// start node: 0,  target node: 5
	fmt.Println("The shortest path: ", shortestPath(edges, 0, 5))
}

func shortestPath(edges [][]int, start int, target int) []int {
	graph := buildGraph(edges)

	queue := []int{start}
	visited := map[int]int{start: -1} // key: node, value: previous node

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == target {
			break
		}

		// look at all its neighbors
		for _, neighbor := range graph[current] {
			// if neighbor is not in visisted
			if _, found := visited[neighbor]; !found {
				visited[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}
	return backtrackPath(visited, target)
}

func buildGraph(edges [][]int) map[int][]int {
	graph := map[int][]int{}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	return graph
}

func backtrackPath(visited map[int]int, end int) []int {
	path := []int{}
	current := end
	for current != -1 {
		path = append(path, current)
		current = visited[current]
	}

	reverse(path)
	return path
}

func reverse(array []int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

/* output
The shortest path:  [0 1 5]
*/
