package main

func lengthOfLongestSubstring(s string) int {
	// longest substring
	res := 0
	// map to hold characters we've seen
	seen := map[byte]bool{}

	// go through string with start, end pointers
	for start, end := 0, 0; end < len(s); end++ {
		// keep removing from seen, as so long as we've already
		// seen this character (no repeated characters)
		for seen[s[end]] {
			delete(seen, s[start])
			start++
		}

		// now we can add the character at the end pointer
		seen[s[end]] = true
		// calculate new longest substring (at this point, we know no repeating)
		if end-start+1 > res {
			res = end - start + 1
		}
	}

	return res
}
