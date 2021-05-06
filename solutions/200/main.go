package main

func numIslands(grid [][]byte) int {
	res := 0

	// go thru grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// check if this is an island
			if grid[i][j] == byte('1') {
				// sink and count it
				sink(&grid, i, j)
				res++
			}
		}
	}

	return res
}

// recursive helper method to sink islands
func sink(grid *[][]byte, i int, j int) {
	// base case (out of bounds, or not island)
	if i < 0 || i == len(*grid) || j < 0 || j == len((*grid)[i]) || (*grid)[i][j] != byte('1') {
		return
	}

	// sink
	(*grid)[i][j] = byte('0')

	// go up,down,right,left
	sink(grid, i+1, j)
	sink(grid, i-1, j)
	sink(grid, i, j+1)
	sink(grid, i, j-1)
}
