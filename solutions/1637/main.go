package main

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	res := 0
	// sort by x coordinates
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	// go through points
	for i := 1; i < len(points); i++ {
		// widest vertical width
		if points[i][0]-points[i-1][0] > res {
			res = points[i][0] - points[i-1][0]
		}
	}

	return res
}
