package main

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	counts, res := map[byte]int{}, 0

	// go thru string s, and keep track of start/end of substring
	for start, end := 0, 0; end < len(s); end++ {
		// increment current character at end
		counts[s[end]]++

		// keep drawing from left side, until we have k distinct
		for len(counts) > k {
			counts[s[start]]--
			if counts[s[start]] == 0 {
				delete(counts, s[start])
			}
			start++
		}

		// check if we have longer substring
		if end-start+1 > res {
			res = end - start + 1
		}
	}

	return res
}
