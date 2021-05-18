package main

func longestConsecutive(nums []int) int {
	res := 0
	// seen set
	seen := map[int]bool{}

	// add all nums to seen
	for _, num := range nums {
		seen[num] = true
	}

	// go back through nums
	for _, num := range nums {
		// make sure we didnt already count the streak
		if !seen[num+1] {
			// walk back through the streak
			start := num - 1
			for seen[start] {
				// decrease start point
				start--
			}

			// check if this streak is bigger than res
			if num-start > res {
				res = num - start
			}
		}
	}

	return res
}
