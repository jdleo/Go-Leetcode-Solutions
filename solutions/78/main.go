package main

import "sort"

func subsets(nums []int) [][]int {
	res := [][]int{}
	// sort nums
	sort.Ints(nums)
	// start backtracking
	backtrack(&res, nums, &[]int{}, 0)
	return res
}

func backtrack(res *[][]int, nums []int, temp *[]int, start int) {
	// add copy of temp to res
	tmp := make([]int, len(*temp))
	copy(tmp, *temp)
	*res = append(*res, tmp[:])

	// go through all possible candidates
	for i := start; i < len(nums); i++ {
		// add to temp
		*temp = append(*temp, nums[i])
		// backtrack
		backtrack(res, nums, temp, i+1)
		// pop from temp
		*temp = (*temp)[:len(*temp)-1]
	}
}
