package main

func groupAnagrams(strs []string) [][]string {
	// anagram buckets
	anagrams := make(map[string][]string)

	// go through strings
	for _, str := range strs {
		// serialize string to anagram form
		a := serialize(str)
		// check if this has a bucket
		if _, ok := anagrams[a]; ok {
			// append to it
			anagrams[a] = append(anagrams[a], str)
		} else {
			// create it
			anagrams[a] = []string{str}
		}
	}

	res := [][]string{}

	// fill res with buckets
	for _, v := range anagrams {
		res = append(res, v)
	}

	return res
}

// helper method to serialize a string into anagram form
func serialize(str string) string {
	// character counts
	counts := make([]byte, 26)
	// fill counts
	for _, char := range str {
		counts[char-'a']++
	}
	// string representation of byte slice
	return string(counts)
}
