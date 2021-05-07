package main

func rob(nums []int) int {
	// edge case
	if len(nums) == 1 {
		return nums[0]
	}

	// pre-compute first two dp spaces
	dp := []int{nums[0], max(nums[0], nums[1])}

	// go through rest of nums, to fill dp
	for i := 2; i < len(nums); i++ {
		// calculate money we can get by robbing or skipping
		rob := nums[i] + dp[i-2]
		skip := dp[i-1]
		// take max, put in dp
		dp = append(dp, max(rob, skip))
	}

	// result is last space in dp array
	return dp[len(dp)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
