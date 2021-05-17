package main

import "math"

func canMakeArithmeticProgression(arr []int) bool {
	min, max, n := math.MaxInt64, math.MinInt64, len(arr)
	seen := map[int]bool{}

	// go thru nums, add to set, and calc min/max
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
		seen[num] = true
	}

	// get diff
	diff := max - min

	// check if evenly divisible by n-1
	if diff%(n-1) != 0 {
		return false
	}

	// divide by n-1
	diff /= (n - 1)

	// go n times
	for i := 0; i < n; i++ {
		// make sure seen contains the min
		if !seen[min] {
			return false
		}

		// add step to min
		min += diff
	}

	// no problems
	return true
}
