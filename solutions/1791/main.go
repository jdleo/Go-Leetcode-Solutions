package main

func findCenter(edges [][]int) int {
	// only one node in multiple edges, so check if we can find that node
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else {
		return edges[0][1]
	}
}
