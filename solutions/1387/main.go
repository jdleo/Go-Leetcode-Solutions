package main

import (
	"sort"
)

func getKth(lo int, hi int, k int) int {
	cache := map[int]int{1: 1}
	powers := make([][]int, hi-lo+1)
	// go from lo to hi
	for i := lo; i < hi+1; i++ {
		// get power val
		powers[i-lo] = []int{i, power(cache, i)}
	}
	// sort using custom comparator
	sort.SliceStable(powers, func(i, j int) bool {
		// sort by power value
		return powers[i][1] < powers[j][1]
	})
	// return kth
	return powers[k-1][0]
}

// helper method to compute power value
func power(cache map[int]int, x int) int {
	// check if not in cache
	if cache[x] == 0 {
		// if x is even then x = x / 2
		// if x is odd then x = 3 * x + 1
		if x&1 == 0 {
			cache[x] = 1 + power(cache, x/2)
		} else {
			cache[x] = 1 + power(cache, 3*x+1)
		}
	}

	return cache[x]
}
