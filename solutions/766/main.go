package main

func isToeplitzMatrix(matrix [][]int) bool {
	// map to hold diagonal values
	diag := map[int]int{}

	// go thru matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			// check if diagonal exists
			if v, ok := diag[i-j]; ok {
				// make sure it matches value
				if matrix[i][j] != v {
					return false
				}
			} else {
				// store it
				diag[i-j] = matrix[i][j]
			}
		}
	}

	// no issues encountered
	return true
}
