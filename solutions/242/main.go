package main

func isAnagram(s string, t string) bool {
	// counts of characters
	counts := make([]int, 26)
	// fill counts
	for _, char := range s {
		counts[char-'a']++
	}
	for _, char := range t {
		counts[char-'a']--
	}

	// make sure balanced
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}

	return true
}
