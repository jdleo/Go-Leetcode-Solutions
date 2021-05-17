package main

func countComponents(n int, edges [][]int) int {
	res := 0
	adj := make([][]int, n) // adjacency list

	// go through edges and build adjacency list
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	// to keep track of nodes explored
	visited := make([]bool, n)

	// go from 0 to n-1
	for i := 0; i < n; i++ {
		// check if we havent visited this yet
		if !visited[i] {
			res++
			// queue for bfs
			queue := []int{i}
			// go while queue is not empty
			for len(queue) > 0 {
				// popleft from queue
				cur := queue[0]
				queue = queue[1:]
				visited[cur] = true

				// add neighbors of cur to queue
				for _, edge := range adj[cur] {
					if !visited[edge] {
						queue = append(queue, edge)
					}
				}

			}
		}
	}

	return res
}
