package main

func isValid(s string) bool {
	// closing tokens, with matching opening tokens
	match := map[rune]rune{')': '(', ']': '[', '}': '{'}
	// stack to hold opening tokens
	stack := []rune{}

	// go through string
	for _, char := range s {
		// check if closing token
		if v, ok := match[char]; ok {
			n := len(stack)
			// check if stack empty, or character in stack is our matching token
			if n == 0 || stack[n-1] != v {
				return false
			}

			// pop from stack
			stack = stack[:n-1]
		} else {
			// just push token to stack
			stack = append(stack, char)
		}
	}

	// stack must be empty (all tokens accounted for)
	return len(stack) == 0
}
