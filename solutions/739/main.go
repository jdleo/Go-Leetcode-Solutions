package main

func dailyTemperatures(T []int) []int {
	// stack to hold temperature indices, and result array
	stack, res := []int{}, make([]int, len(T))

	// go through temperatures
	for i, t := range T {
		// keep drawing from stack, as so long as current temp is warmer
		for len(stack) > 0 && T[stack[len(stack)-1]] < t {
			// add distance to res of current idx - idx of colder temp
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			// pop from stack
			stack = stack[:len(stack)-1]
		}
		// add index to stack
		stack = append(stack, i)
	}

	return res
}
