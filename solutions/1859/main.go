package main

import "strings"

func sortSentence(s string) string {
	// split by space
	words := strings.Split(s, " ")
	res := make([]string, len(words))

	// go through words
	for _, word := range words {
		// get 1-index
		index := word[len(word)-1] - '1'
		// insert at index
		res[index] = word[:len(word)-1]
	}

	// join back
	return strings.Join(res, " ")
}
