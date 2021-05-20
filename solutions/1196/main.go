package main

import "sort"

func maxNumberOfApples(arr []int) int {
	weight, res := 5000, 0
	// sort ascending
	sort.Ints(arr)

	// go thru apples
	for _, apple := range arr {
		// check if exceeds weight
		if weight-apple < 0 {
			break
		}
		// add apple
		weight -= apple
		res++
	}

	return res
}
