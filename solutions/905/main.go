package main

func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1

	// go until ptrs meet
	for left < right {
		if nums[left]&1 == 0 {
			// left is even and in correct spot, keep going
			left++
		} else if nums[right]&1 == 1 {
			// right is odd and in correct spot, keep going
			right--
		} else {
			// they need to be swapped
			nums[left], nums[right] = nums[right], nums[left]
			left, right = left+1, right-1
		}
	}

	return nums
}
