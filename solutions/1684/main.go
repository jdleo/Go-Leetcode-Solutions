package main

func countConsistentStrings(allowed string, words []string) int {
	res := 0

	// allowed character set
	valid := make([]bool, 26)
	for _, char := range allowed {
		valid[char-'a'] = true
	}

	// go through words
OUTER:
	for _, word := range words {

		// go through word
		for _, char := range word {
			// check if inconsistent
			if !valid[char-'a'] {
				// continue
				continue OUTER
			}
		}

		// no issues
		res++
	}

	return res
}
