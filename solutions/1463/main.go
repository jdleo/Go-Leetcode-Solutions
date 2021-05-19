package main

import "math"

type Key struct {
	i  int
	j1 int
	j2 int
}

func cherryPickup(grid [][]int) int {
	memo := map[Key]int{}
	return backtrack(memo, grid, 0, 0, len(grid[0])-1)
}

// helper method to recursively backtrack
func backtrack(memo map[Key]int, grid [][]int, i, j1, j2 int) int {
	// bounds check
	if j1 < 0 || j1 == len(grid[0]) || j2 < 0 || j2 == len(grid[0]) {
		return math.MinInt64
	} else if i == len(grid) {
		return 0
	}

	// check if in memo
	key := Key{i, j1, j2}
	if val, ok := memo[key]; ok {
		return val
	}

	res := 0

	// go all possible directions
	for _, dj1 := range []int{-1, 0, 1} {
		for _, dj2 := range []int{-1, 0, 1} {
			res = max(res, backtrack(memo, grid, i+1, j1+dj1, j2+dj2))
		}
	}

	// pick these cherries
	if j1 != j2 {
		res += grid[i][j1] + grid[i][j2]
	} else {
		res += grid[i][j1]
	}

	memo[key] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
