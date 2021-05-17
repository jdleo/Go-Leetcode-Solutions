package main

import "sort"

func heightChecker(heights []int) int {
	// sorted version of heights
	actual := make([]int, len(heights))
	copy(actual, heights)
	sort.Ints(actual)

	res := 0

	for i, height := range heights {
		// check if differs
		if height != actual[i] {
			res++
		}
	}

	return res
}
