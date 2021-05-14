package main

import "strings"

func replaceDigits(s string) string {
	// result string builder, allocate s length to it
	var res strings.Builder
	res.Grow(len(s))

	// go through s
	for i, char := range s {
		// check if this char is a digit
		if char < 'a' {
			// write the last byte + current byte (digit)
			res.WriteByte(s[i-1] + byte(char-'0'))
		} else {
			// just write this rune
			res.WriteRune(char)
		}
	}

	return res.String()
}
