package main

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}

	for i, row := range res {
		for j := range row {
			res[i][j] = min(rowSum[i], colSum[j])
			rowSum[i] -= res[i][j]
			colSum[j] -= res[i][j]
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
