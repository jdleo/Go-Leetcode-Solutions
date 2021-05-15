package main

func sortString(s string) string {
	counts := make([]int, 26)
	res := []byte{}
	n := len(s)

	// fill counts
	for _, char := range s {
		counts[char-'a']++
	}

	// go until we've picked all characters
	for n > 0 {
		// go through smallest characters (0 -> 25)
		for i := 0; i < 26; i++ {
			// check if we have a character to pick
			if counts[i] > 0 {
				// pick this character
				res = append(res, byte('a'+i))
				counts[i]--
				n--
			}
		}
		// go through largest characters (25 -> 0)
		for i := 25; i > -1; i-- {
			// check if we have a character to pick
			if counts[i] > 0 {
				// pick this character
				res = append(res, byte('a'+i))
				counts[i]--
				n--
			}
		}
	}

	// return byte slice as string
	return string(res)
}
