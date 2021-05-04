package main

func minPathSum(grid [][]int) int {
	// fill first col, and first row w running sum (only one direction here)
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] += grid[0][j-1]
	}

	// go thru rest of grid
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {

		}
	}

	// result is in last cell
	return grid[len(grid)-1][len(grid[0])-1]
}
