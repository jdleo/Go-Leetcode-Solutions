package main

func luckyNumbers(matrix [][]int) []int {
	res := []int{}
	// lucky number is minRow maxCol
	rowMin := make([]int, len(matrix))
	colMax := make([]int, len(matrix[0]))

	// go thru matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// set row min
			if rowMin[i] == 0 || matrix[i][j] < rowMin[i] {
				rowMin[i] = matrix[i][j]
			}

			// set col max
			if matrix[i][j] > colMax[j] {
				colMax[j] = matrix[i][j]
			}
		}
	}

	// go thru matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// check if rowmin and colmax
			if matrix[i][j] == rowMin[i] && matrix[i][j] == colMax[j] {
				res = append(res, matrix[i][j])
			}
		}
	}

	return res
}
