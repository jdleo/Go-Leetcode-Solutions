package main

func productExceptSelf(nums []int) []int {
	// left and right running products
	left, right := 1, 1
	// result array filled with 1s
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = 1
	}

	// iterate thru nums length
	for i := 0; i < len(nums); i++ {
		// multiply into to res the current left and right products
		res[i] *= left
		res[len(res)-i-1] *= right
		// multiply into left and right products
		left *= nums[i]
		right *= nums[len(res)-i-1]
	}

	return res
}
