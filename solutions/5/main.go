package main

func longestPalindrome(s string) string {
	// current longest
	start, end := 0, 0

	// go thru length of s
	for i := 0; i < len(s); i++ {
		// calculate longest odd/even palindrom centered at this index
		oddStart, oddEnd := expand(s, i, i)
		evenStart, evenEnd := expand(s, i, i+1)
		// set new max
		if oddEnd-oddStart > end-start {
			start, end = oddStart, oddEnd
		}
		if evenEnd-evenStart > end-start {
			start, end = evenStart, evenEnd
		}
	}

	// return string slice of start:end
	return s[start : end+1]
}

func expand(s string, i, j int) (int, int) {
	// keep going, as so long as this isn't out of bounds, and IS palindrome
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i, j = i-1, j+1
	}

	// this is our palindrome ending
	return i + 1, j - 1
}
