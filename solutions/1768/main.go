package main

func mergeAlternately(word1 string, word2 string) string {
	res := make([]byte, len(word1)+len(word2))
	i, j := 0, 0

	// go while word1 and word2 have chars
	for i < len(word1) && j < len(word2) {
		// interweave
		res[i+j] = word1[i]
		res[i+j+1] = word2[j]
		i, j = i+1, j+1
	}

	// go while word1 still has chars
	for i < len(word1) {
		res[i+j] = word1[i]
		i++
	}

	// go while word2 still has chars
	for j < len(word2) {
		res[i+j] = word2[j]
		j++
	}

	// byte slice as string
	return string(res)
}
