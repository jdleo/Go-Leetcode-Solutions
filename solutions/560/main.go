package main

func subarraySum(nums []int, k int) int {
	res, sum := 0, 0

	// prefix sum map
	cache := map[int]int{0: 1}

	// go through nums
	for _, num := range nums {
		// add num to current sum
		sum += num
		// check if sum-k exists in our cache
		if v, ok := cache[sum-k]; ok {
			// add to res
			res += v
		}
		// also add this sum to prefix sums
		cache[sum]++
	}

	return res
}
