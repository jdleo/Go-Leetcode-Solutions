package main

func lengthOfLongestSubstring(s string) int {
	// longest substring
	res := 0
	// map to hold characters we've seen
	seen := map[byte]bool{}

	// go through string with start, end pointers
	for start, end := 0, 0; end < len(s); end++ {
		// keep removing from seen, and moving start forward
		// as so long as we've already seen this char
		for seen[s[end]] {
			delete(seen, s[start])
			start++
		}

		// ok, now we can add character at end pointer
		seen[s[end]] = true
		// set new max length
		if end-start+1 > res {
			res = end - start + 1
		}
	}

	return res
}
