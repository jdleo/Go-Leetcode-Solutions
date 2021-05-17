package main

func uniqueOccurrences(arr []int) bool {
	counts := map[int]int{}
	seen := map[int]bool{}

	// fill frequencies of each num
	for _, num := range arr {
		counts[num]++
	}

	// go through freqs
	for _, freq := range counts {
		// non-unique freq
		if seen[freq] {
			return false
		}

		seen[freq] = true
	}

	return true
}
