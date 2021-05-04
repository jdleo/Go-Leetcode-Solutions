package main

func countSubstrings(s string) int {
	res := 0

	// go through each index in s
	for i := 0; i < len(s); i++ {
		// count palindromic substrings of both even and odd length
		res += expand(s, i, i+1)
		res += expand(s, i, i)
	}

	return res
}

func expand(s string, left int, right int) int {
	res := 0

	// keep expanding as so long as the characters are equal
	for left >= 0 && right < len(s) && s[left] == s[right] {
		// count substring
		res += 1
		// expand
		left, right = left-1, right+1
	}

	return res
}
