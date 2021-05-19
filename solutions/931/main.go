package main

import "math"

func minFallingPathSum(matrix [][]int) int {
	// go through matrix
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// take this cell, plus the minimum of the 3 top
			curMin := matrix[i-1][j]
			// check if not out of bounds
			if j > 0 {
				curMin = min(curMin, matrix[i-1][j-1])
			}
			if j < len(matrix[0])-1 {
				curMin = min(curMin, matrix[i-1][j+1])
			}
			// add the min + current
			matrix[i][j] += curMin
		}
	}

	res := math.MaxInt64

	// get minimum of last row
	for j := 0; j < len(matrix[0]); j++ {
		res = min(res, matrix[len(matrix)-1][j])
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
