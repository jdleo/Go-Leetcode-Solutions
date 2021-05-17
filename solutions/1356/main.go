package main

import "sort"

func sortByBits(arr []int) []int {
	cache := map[int]int{0: 0, 1: 1}
	// sort numbers by their bit value, ascending if equal
	sort.Slice(arr, func(i, j int) bool {
		a, b := bits(cache, arr[i]), bits(cache, arr[j])

		if a == b {
			return arr[i] < arr[j]
		}

		return a < b
	})
	return arr
}

// recursive helper method to calculate set bits in a number
func bits(cache map[int]int, x int) int {
	// check if not in map
	if _, ok := cache[x]; !ok {
		// calculate bits after left shifting
		cache[x] = x&1 + bits(cache, x>>1)
	}

	return cache[x]
}
