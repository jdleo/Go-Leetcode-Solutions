package main

func freqAlphabets(s string) string {
	res := []byte{}

	// go thru string backwards
	for i := len(s) - 1; i > -1; i-- {
		// check if pound
		if s[i] == '#' {
			// get the next two bytes, and concat them, then add to front
			digit := (s[i-2]-'0')*10 + (s[i-1] - '0') - 1
			char := 'a' + digit
			res = append([]byte{char}, res...)
			i -= 2
		} else {
			// just add the character from 'a' 1-indexed, to front
			char := 'a' + (s[i] - '1')
			res = append([]byte{char}, res...)
		}
	}

	// byte slice as string
	return string(res)
}
