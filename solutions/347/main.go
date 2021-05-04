package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	// map to hold number frequencies
	counts := map[int]int{}
	// fill frequencies
	for _, num := range nums {
		counts[num]++
	}

	// keys of frequencies (values)
	keys := []int{}
	for k := range counts {
		keys = append(keys, k)
	}

	// sort keys based on frequency
	sort.Slice(keys, func(i int, j int) bool {
		return counts[keys[i]] > counts[keys[j]]
	})

	// return k most frequent
	return keys[:k]
}
