package main

func maxProduct(nums []int) int {
	// minimum/maximum current running product, and res
	minProd, maxProd, res := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		// compute min/max of either starting new subarray, or continuing subarray
		x := min(nums[i], min(nums[i]*minProd, nums[i]*maxProd))
		y := max(nums[i], max(nums[i]*minProd, nums[i]*maxProd))

		// set minprod, maxprod
		minProd, maxProd = x, y

		// set new current max product subarray
		res = max(res, maxProd)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
