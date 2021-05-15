package main

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0

	// go thru pointers
	for i := 1; i < len(points); i++ {
		// max of the delta between x and y coords
		res += max(abs(points[i][0]-points[i-1][0]), abs(points[i][1]-points[i-1][1]))
	}

	return res
}

// helper methods below

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
