package main

func balancedStringSplit(s string) int {
	res, bal := 0, 0

	// go through string
	for _, char := range s {
		// adjust balance based on L/R characters
		switch char {
		case 'L':
			bal++
		case 'R':
			bal--
		}

		// check if evenly balanced
		if bal == 0 {
			res++
		}
	}

	return res
}
