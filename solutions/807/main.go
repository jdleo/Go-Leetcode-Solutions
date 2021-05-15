package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	res := 0
	// get dimensions
	m, n := len(grid), len(grid[0])
	// max per each row, per each col
	rowMax := make([]int, m)
	colMax := make([]int, n)

	// go thru grid and calculate rowmax and colmax
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rowMax[i] = max(rowMax[i], grid[i][j])
			colMax[j] = max(colMax[j], grid[i][j])
		}
	}

	// go back thru grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// we can increase current cell UP to the min of the maxes
			res += min(rowMax[i], colMax[j]) - grid[i][j]
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
