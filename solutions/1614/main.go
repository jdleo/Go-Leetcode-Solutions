package main

func maxDepth(s string) int {
	// current depth, and max depth we've seen
	depth, res := 0, 0

	// go through s
	for _, char := range s {
		// check if open parentheses (going deeper)
		if char == '(' {
			depth++
			// set new max, if needed
			if depth > res {
				res = depth
			}
		} else if char == ')' {
			// decrease depth
			depth--
		}
	}

	return res
}
