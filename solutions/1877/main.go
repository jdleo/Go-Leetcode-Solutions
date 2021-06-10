package main

import "sort"

func minPairSum(nums []int) int {
	var res int
	// first, sort
	sort.Ints(nums)

	// go up until middle
	for i := 0; i < len(nums)/2; i++ {
		// sum first and last
		sum := nums[i] + nums[len(nums)-i-1]
		// set new max
		if sum > res {
			res = sum
		}
	}

	return res
}
