package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	res := []int{}

	// array[i] means that node i has indegree array[i]
	indegree := make([]bool, n)

	// go through edges
	for _, edge := range edges {
		// mark indegree
		indegree[edge[1]] = true
	}

	for i := 0; i < n; i++ {
		// check if this node has indegree of 0
		if !indegree[i] {
			res = append(res, i)
		}
	}

	return res
}
