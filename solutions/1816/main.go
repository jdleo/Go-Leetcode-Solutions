package main

func truncateSentence(s string, k int) string {
	res := []byte{}

	// go through string
	for _, char := range s {
		// check if this is a space
		if char == ' ' {
			// decrease k
			k--
			// check if k is at 0 (we have all our words)
			if k == 0 {
				return string(res)
			}
			// push space
			res = append(res, byte(char))
		} else {
			// push char
			res = append(res, byte(char))
		}
	}

	// return byte slice as string
	return string(res)
}
