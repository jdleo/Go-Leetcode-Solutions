package main

func countSquares(matrix [][]int) int {
	res := 0

	// go thru matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// make sure this is a 1, and not an edge
			if matrix[i][j] == 1 && i > 0 && j > 0 {
				// get min of left, top, top left (to see if we can continue a square)
				matrix[i][j] = 1 + min(matrix[i-1][j], min(matrix[i][j-1], matrix[i-1][j-1]))
			}
			// add this current cell to res
			res += matrix[i][j]
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
