package main

import "strings"

func uniqueMorseRepresentations(words []string) int {
	// mapping for char to morse
	morse := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....",
		"..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.",
		"--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-",
		"-.--", "--..",
	}

	// set for keeping track of words
	set := map[string]struct{}{}

	// go through words
	for _, word := range words {
		var cur strings.Builder
		// go through character in word
		for _, char := range word {
			// add morse encoding to cur
			cur.WriteString(morse[char-'a'])
		}
		// add current word to set
		set[cur.String()] = struct{}{}
	}

	return len(set)
}
