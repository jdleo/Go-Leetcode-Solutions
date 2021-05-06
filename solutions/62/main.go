package main

func uniquePaths(m int, n int) int {
	// create dp array
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// first row 1s
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	// first col 1s
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// go thru rest
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// unique paths is paths from top, and left
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// result is in bottom right cell of dp
	return dp[m-1][n-1]
}
