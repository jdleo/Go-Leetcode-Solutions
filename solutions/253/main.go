package main

import "sort"

func minMeetingRooms(intervals [][]int) int {
	var points []point
	for _, interval := range intervals {
		points = append(points, point{interval[0], true}, point{interval[1], false})
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].time == points[j].time {
			return !points[i].isStart
		}
		return points[i].time < points[j].time
	})

	var rooms, max int
	for _, pt := range points {
		if pt.isStart {
			rooms++
		} else {
			rooms--
		}
		if rooms > max {
			max = rooms
		}
	}
	return max
}
