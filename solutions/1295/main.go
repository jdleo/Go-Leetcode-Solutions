package main

import "math"

func findNumbers(nums []int) int {
	res := 0

	// go thru nums
	for _, num := range nums {
		// to get num of digits, its floor(log10(n) + 1)
		if int(math.Floor(math.Log10(float64(num))+1))%2 == 0 {
			res++
		}
	}

	return res
}
