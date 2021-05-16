package main

func findAndReplacePattern(words []string, pattern string) []string {
	res := []string{}
	pattern_ := normalize(pattern)

	// go thru words
	for _, word := range words {
		// check if normalized version of this word == pattern
		if normalize(word) == pattern_ {
			res = append(res, word)
		}
	}

	return res
}

// helper function to normalize string
func normalize(word string) string {
	// map, and pointer of chars to first occurrence
	seen, res := map[byte]byte{}, []byte{}
	var ptr byte = '0'

	// go through word
	for i := 0; i < len(word); i++ {
		// check if not in seen
		if seen[word[i]] == 0 {
			seen[word[i]] = ptr
			ptr++
		}

		// add its mapping
		res = append(res, seen[word[i]])
	}

	return string(res)
}
