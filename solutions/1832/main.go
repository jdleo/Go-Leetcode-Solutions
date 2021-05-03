package main

func checkIfPangram(sentence string) bool {
	seen := make([]bool, 26)
	res := 26

	for _, char := range sentence {
		if !seen[char-'a'] {
			res--
		}
		seen[char-'a'] = true
	}

	return res == 0
}
