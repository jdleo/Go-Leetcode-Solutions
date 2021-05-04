package main

func countBits(num int) []int {
	// dp array with size of num+1
	dp := make([]int, num+1)

	// check if greater than 0
	if num > 0 {
		// precalculate 1
		dp[1] = 1
	}

	// go from 2 to num
	for i := 2; i <= num; i++ {
		// number has same num of bits as num/2, and then just add 1 if its odd (an extra bit)
		dp[i] = dp[i>>1] + i&1
	}

	// return dp
	return dp
}
