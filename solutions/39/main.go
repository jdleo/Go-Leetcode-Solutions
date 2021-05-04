package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	// sort candidates ascending
	sort.Ints(candidates)
	// result array
	res := [][]int{}
	// start backtracking
	backtrack(&res, candidates, &[]int{}, target, 0)
	return res
}

func backtrack(res *[][]int, candidates []int, path *[]int, target int, start int) {
	// check if we either found a combination sum, or went over
	if target == 0 {
		// add copy of path to res
		temp := make([]int, len(*path))
		copy(temp, *path)
		*res = append(*res, temp[:])
	} else if target < 0 {
		// break
		return
	}

	// go through all possible candidates
	for i := start; i < len(candidates); i++ {
		// take this candidate
		*path = append(*path, candidates[i])
		// backtrack
		backtrack(res, candidates, path, target-candidates[i], i)
		// pop
		*path = (*path)[:len(*path)-1]
	}
}
