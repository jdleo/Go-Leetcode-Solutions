package main

func largestUniqueNumber(A []int) int {
	// counts for each num
	counts := map[int]int{}
	// count each num
	for _, num := range A {
		counts[num]++
	}

	res := -1

	// go thru counts
	for num, count := range counts {
		// check if unique
		if count == 1 {
			// check if bigger
			if num > res {
				res = num
			}
		}
	}

	return res
}
