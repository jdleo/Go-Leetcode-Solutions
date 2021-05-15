package main

func minAddToMakeValid(s string) int {
	res, bal := 0, 0

	// go through string
	for _, char := range s {
		switch char {
		case '(':
			// open up new parantheses
			bal++
		default:
			// check if we have a matching parantheses to close
			if bal > 0 {
				bal--
			} else {
				// we didn't, so we had to add
				res++
			}
		}
	}

	// res + bal (this means we have to still close some)
	return res + bal
}
