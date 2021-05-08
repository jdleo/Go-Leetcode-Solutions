package main

type Pair struct {
	target, cur int
}

func findTargetSumWays(nums []int, target int) int {
	// cache for memoization
	cache := map[Pair]int{}
	// backtracking helper function
	return backtrack(cache, nums, target, 0)
}

func backtrack(cache map[Pair]int, nums []int, target int, cur int) int {
	// base case (reached end)
	if cur == len(nums) {
		// check if we reached our target sum
		if target == 0 {
			return 1
		}
		return 0
	}

	// key for memo
	key := Pair{target, cur}

	// check if this is cached
	if v, ok := cache[key]; ok {
		return v
	}

	// calculate left and right subtrees
	// and either taking the - or the + of current num in nums
	left := backtrack(cache, nums, target-nums[cur], cur+1)
	right := backtrack(cache, nums, target+nums[cur], cur+1)
	cache[key] = left + right
	return left + right
}
