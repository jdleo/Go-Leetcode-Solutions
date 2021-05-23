package main

import (
	"strings"
)

func toGoatLatin(sentence string) string {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true,
		'o': true, 'u': true, 'A': true,
		'E': true, 'I': true, 'O': true,
		'U': true,
	}

	// split sentence
	words := strings.Split(sentence, " ")

	// go thru words
	for i, word := range words {
		// check if doesn't start with vowel
		if !vowels[word[0]] {
			// take first character and put it in back
			word = word[1:] + word[:1]
		}

		// add ma + i number of 'a's
		var newWord strings.Builder
		newWord.WriteString(word + "ma")
		for j := 0; j < i+1; j++ {
			newWord.WriteByte('a')
		}
		words[i] = newWord.String()
	}

	return strings.Join(words, " ")
}
