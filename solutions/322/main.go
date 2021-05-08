package main

const inf = 64 << 1

func coinChange(coins []int, amount int) int {
	// make dp array of size amount+1 (because we want to account for 0 coins too)
	dp := make([]int, amount+1)
	// fill the rest of dp with infinites
	for i := 1; i <= amount; i++ {
		dp[i] = inf
	}

	// go through all possible amounts
	for i := 1; i <= amount; i++ {
		// go through all possible coins to use
		for _, coin := range coins {
			// check if we can use this coin to contribute to amount
			if i-coin >= 0 {
				// either don't take this coin, or do
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// check if last value is infinite (could never make change)
	if dp[len(dp)-1] == inf {
		return -1
	}

	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
