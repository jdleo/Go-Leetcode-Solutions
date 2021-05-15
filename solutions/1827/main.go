package main

func minOperations(nums []int) int {
	res := 0

	for i := 1; i < len(nums); i++ {
		// check if this is non-increasing
		if nums[i] <= nums[i-1] {
			// we have to increase current num to make it strictly increasing
			res += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}

	return res
}
