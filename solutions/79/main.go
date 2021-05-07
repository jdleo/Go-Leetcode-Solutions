package main

func exist(board [][]byte, word string) bool {
	// go thru board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			// try to backtrack on this cell
			if backtrack(board, word, 0, i, j) {
				return true
			}
		}
	}
	// never found
	return false
}

// helper function to backtrack thru grid
func backtrack(board [][]byte, word string, cur, i, j int) bool {
	// base cases
	if cur == len(word) {
		return true
	} else if i < 0 || i == len(board) || j < 0 || j == len(board[i]) {
		return false
	} else if word[cur] != board[i][j] {
		return false
	}

	// take character out of search space
	tmp := board[i][j]
	board[i][j] = '#'

	// go up,down,right,left
	res := (backtrack(board, word, cur+1, i+1, j) ||
		backtrack(board, word, cur+1, i-1, j) ||
		backtrack(board, word, cur+1, i, j+1) ||
		backtrack(board, word, cur+1, i, j-1))

	// put board piece back into search space
	board[i][j] = tmp

	return res
}
