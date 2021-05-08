package main

import "math"

func jump(nums []int) int {
	// dp array
	dp := make([]int, len(nums))
	// fill all with infinity (except first)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt64
	}

	// build up dp (min jumps to get to cell i)
	for i := 1; i < len(dp); i++ {
		// all possible indices we could have came from
		for j := 0; j < i; j++ {
			// check if we could have come from this index
			if nums[j] >= i-j {
				// update min jumps for this current index (this is +1 jump)
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	// answer is in last cell
	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
