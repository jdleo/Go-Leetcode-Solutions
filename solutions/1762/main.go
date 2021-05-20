package main

func findBuildings(heights []int) []int {
	res, maxes, curMax := []int{}, make([]int, len(heights)), -1

	for i := len(maxes) - 1; i > -1; i-- {
		// set current max from right, and new curmax
		maxes[i], curMax = curMax, max(curMax, heights[i])
	}

	// go back thru heights once more
	for i := 0; i < len(heights); i++ {
		// check if building is taller than tallest from right
		if heights[i] > maxes[i] {
			res = append(res, i)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
