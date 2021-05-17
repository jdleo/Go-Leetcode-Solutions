package main

func sortArrayByParityII(nums []int) []int {
	left, right := 0, 1

	for left < len(nums) {
		// check if left is out of place
		if nums[left]&1 == 1 {
			// swap with whats at right, and move right forward
			nums[left], nums[right], right = nums[right], nums[left], right+2
		} else {
			left += 2
		}
	}

	return nums
}
