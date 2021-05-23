package main

func islandPerimeter(grid [][]int) int {
	res := 0

	// go thru grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// check if this is an island
			if grid[i][j] == 1 {
				res += 4

				// check if left was already counted
				if j > 0 && grid[i][j-1] == 1 {
					res -= 2
				}

				// check if top was already counted
				if i > 0 && grid[i-1][j] == 1 {
					res -= 2
				}
			}
		}
	}

	return res
}
