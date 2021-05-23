package main

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))

	// go until left and right meet
	for i, left, right := 0, 0, len(nums)-1; left <= right; i++ {
		// take bigger number squared
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[i] = nums[left] * nums[left]
			left++
		} else {
			res[i] = nums[right] * nums[right]
			right--
		}
	}

	// reverse res
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		res[left], res[right] = res[right], res[left]
	}

	return res
}
