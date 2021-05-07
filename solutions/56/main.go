package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// sort intervals by starts
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// stack to hold merged intervals (first one is here to start)
	stack := [][]int{intervals[0]}

	// go through rest of intervals
	for i := 1; i < len(intervals); i++ {
		n := len(stack) - 1
		// is current start, less or equal to last end?
		if intervals[i][0] <= stack[n][1] {
			// we want to merge them
			start := stack[n][0]
			end := max(intervals[i][1], stack[n][1])
			// replace last (merge) with new interval
			stack[n] = []int{start, end}
		} else {
			// just add this interval, doesnt overlap
			stack = append(stack, intervals[i])
		}
	}

	// stack will have merged intervals
	return stack
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
