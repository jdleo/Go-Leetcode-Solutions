package main

import "math"

func findDisappearedNumbers(nums []int) []int {
	// go through nums
	for _, n := range nums {
		n := abs(n)
		if nums[n-1] > 0 {
			nums[n-1] *= -1
		}
	}

	res := []int{}

	// go through nums again
	for i, n := range nums {
		// check if this number was never marked
		if n > 0 {
			// this means the index for it was never present
			res = append(res, i+1)
		}
	}

	return res
}

func abs(num int) int {
	return int(math.Abs(float64(num)))
}
