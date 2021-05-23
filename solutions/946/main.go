package main

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i := 0

	// go through push sequence
	for _, num := range pushed {
		// push to stack
		stack = append(stack, num)

		// keep drawing from stack, as long as we can pop
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			// pop from stack, increment ptr
			stack = stack[:len(stack)-1]
			i++
		}
	}

	// are all accounted for?
	return len(stack) == 0
}
