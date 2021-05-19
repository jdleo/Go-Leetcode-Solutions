package main

func countLetters(s string) int {
	res := 0
	same := 1 // initialize to one, since first character we see is valid

	for i, char := range s {
		if i > 0 {
			if s[i-1] == byte(char) {
				same++
			} else {
				same = 1
			}
		}

		// add unique substring to res
		res += same
	}

	return res
}
