package main

func maxSubArray(nums []int) int {
	res := nums[0]
	// maximum subarray sums at each index
	maxSums := make([]int, len(nums))
	maxSums[0] = nums[0]

	// go through nums
	for i := 1; i < len(nums); i++ {
		// the max sum at this index, is the max sum of
		// either 1) continuing sum, starting new sum
		maxSums[i] = max(maxSums[i-1]+nums[i], nums[i])
		// new max
		res = max(res, maxSums[i])
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
