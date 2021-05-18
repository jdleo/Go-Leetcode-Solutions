package main

func characterReplacement(s string, k int) int {
	counts := make([]int, 26)
	maxCount, maxLength := 0, 0

	for start, end := 0, 0; end < len(s); end++ {
		// increment count for this character
		counts[s[end]-'A']++
		// set new max count
		maxCount = max(maxCount, counts[s[end]-'A'])
		// while the rest of characters - max count > k
		for end-start+1-maxCount > k {
			// remove from start n increment it
			counts[s[start]-'A']--
			start++
		}
		// set new max length
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
