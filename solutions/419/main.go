package main

func countBattleships(board [][]byte) int {
	res, m, n := 0, len(board), len(board[0])

	// recursive, dfs helper method to sink battleships
	var sink func(i, j int)
	sink = func(i, j int) {
		// base case : out of bounds, or not battleship
		if i < 0 || i == m || j < 0 || j == n || board[i][j] != 'X' {
			return
		}

		// sink
		board[i][j] = '.'

		// go up,down,right,left
		sink(i+1, j)
		sink(i-1, j)
		sink(i, j+1)
		sink(i, j-1)
	}

	// go thru grid
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// check if battle ship
			if board[i][j] == 'X' {
				// sink
				sink(i, j)
				res++
			}
		}
	}

	return res
}
