package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	// go through words
	for i, word := range words {
		// get byte slice
		slice := []byte(word)

		// reverse
		for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}

		// put back in
		words[i] = string(slice)
	}

	// join back
	return strings.Join(words, " ")
}
