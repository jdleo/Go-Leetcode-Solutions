package main

func rotate(matrix [][]int) {
	// first, go through and swap individual rows
	for i := 0; i < len(matrix)/2; i++ {
		// swap row in front with row in back
		matrix[i], matrix[len(matrix)-i-1] = matrix[len(matrix)-i-1], matrix[i]
	}

	// swap mirrored cells
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			// swap cells
			matrix[i], matrix[j] = matrix[j], matrix[i]
		}
	}
}
