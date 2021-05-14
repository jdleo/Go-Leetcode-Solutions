package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	// byte slices for words
	slice1 := []byte{}
	slice2 := []byte{}

	// fill slice1
	for _, word := range word1 {
		slice1 = append(slice1, []byte(word)...)
	}

	// fill slice2
	for _, word := range word2 {
		slice2 = append(slice2, []byte(word)...)
	}

	// check if string slices are equivalent
	return string(slice1) == string(slice2)
}
