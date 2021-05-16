package main

func diStringMatch(s string) []int {
	left, right := 0, len(s)
	res := make([]int, len(s)+1)

	// go thru string
	for i, char := range s {
		// check if increasing
		if char == 'I' {
			res[i] = left
			left++
		} else {
			res[i] = right
			right--
		}
	}

	// set the last one (can be either)
	res[len(s)] = left
	return res
}
