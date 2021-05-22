package main

func maxAreaOfIsland(grid [][]int) int {
	res := 0

	// go thru grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// set new max area, if need be, and sink island
			area := sinkAndGetArea(grid, i, j)
			if area > res {
				res = area
			}
		}
	}

	return res
}

// helper method to 'sink' island, and return area of it
func sinkAndGetArea(grid [][]int, i, j int) int {
	// bounds check
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[i]) || grid[i][j] != 1 {
		return 0
	}

	// sink island, and count for this area
	area := 1
	grid[i][j] = 0

	// all possible directions
	area += sinkAndGetArea(grid, i+1, j)
	area += sinkAndGetArea(grid, i-1, j)
	area += sinkAndGetArea(grid, i, j+1)
	area += sinkAndGetArea(grid, i, j-1)

	return area
}
