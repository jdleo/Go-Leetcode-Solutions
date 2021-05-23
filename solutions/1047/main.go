package main

func removeDuplicates(s string) string {
	stack := []byte{}

	// go thru bytes in s
	for i := 0; i < len(s); i++ {
		// check if stack has bytes, and this byte is equal to last
		if len(stack) != 0 && s[i] == stack[len(stack)-1] {
			// pop
			stack = stack[:len(stack)-1]
		} else {
			// add
			stack = append(stack, s[i])
		}
	}

	// return byte stack as string
	return string(stack)
}
