package main

func firstUniqChar(s string) int {
	// counts of lowercase characters
	counts := make([]int, 26)

	for _, char := range s {
		// increment freq
		counts[char-'a']++
	}

	// go back thru string
	for i, char := range s {
		// see if unique
		if counts[char-'a'] == 1 {
			return i
		}
	}

	// doesnt exist
	return -1
}
