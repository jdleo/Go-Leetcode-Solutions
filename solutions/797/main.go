package main

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	dfs(graph, &res, &[]int{}, 0)
	return res
}

// helper method for dfs
func dfs(graph [][]int, res *[][]int, path *[]int, cur int) {
	// add this node to path
	*path = append(*path, cur)
	// base case : reached end
	if cur == len(graph)-1 {
		// push copy of path to res
		temp := make([]int, len(*path))
		copy(temp, *path)
		*res = append(*res, temp[:])
	} else {
		// go to all neighbors
		for _, neighbor := range graph[cur] {
			// visit
			dfs(graph, res, path, neighbor)
		}
	}
	// remove this node from path
	*path = (*path)[:len(*path)-1]
}
