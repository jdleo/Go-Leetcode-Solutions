package main

func diagonalSum(mat [][]int) int {
	res, n := 0, len(mat)

	// go thru matrix length
	for i := 0; i < n; i++ {
		// add both diagonals
		res += mat[i][i] + mat[n-i-1][i]
	}

	// check if odd (that means we double counted middle)
	if n&1 == 1 {
		return res - mat[n/2][n/2]
	}
	return res
}
