package main

func missingNumber(nums []int) int {
	// set res to be length of nums
	res := len(nums)

	// go thru nums
	for i, num := range nums {
		// xor res with index and num
		res ^= i ^ num
	}

	return res
}
