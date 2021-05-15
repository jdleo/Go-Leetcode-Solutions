package main

func halvesAreAlike(s string) bool {
	// vowel map
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	// balance
	bal := 0

	// go up to half of s
	for i := 0; i < len(s)/2; i++ {
		// first half, increment if vowel
		if vowels[s[i]] {
			bal++
		}
		// second half, decrement if vowel
		if vowels[s[len(s)/2+i]] {
			bal--
		}
	}

	// check if balanced vowels
	return bal == 0
}
