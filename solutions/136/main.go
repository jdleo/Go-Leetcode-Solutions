package main

func singleNumber(nums []int) int {
	res := 0

	// go thru nums
	for _, num := range nums {
		// xor with res (because num xor num = 0. will leave single one stranded)
		res ^= num
	}

	return res
}
