package main

import "math"

func calculateTime(keyboard string, word string) int {
	res, last, indices := 0, 0, make([]int, 26)

	// figure out each keyboard character position
	for i, char := range keyboard {
		indices[char-'a'] = i
	}

	// go through word
	for _, char := range word {
		// add distance from last to current
		res += int(math.Abs(float64(last - indices[char-'a'])))
		// set last
		last = indices[char-'a']
	}

	return res
}
