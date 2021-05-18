package main

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	// sort intervals by start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// go thru intervals
	for i := 1; i < len(intervals); i++ {
		// check if overlaps with last
		if intervals[i-1][1] > intervals[i][0] {
			return false
		}
	}

	// no issues
	return true
}
