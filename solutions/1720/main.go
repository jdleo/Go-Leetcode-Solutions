package main

func decode(encoded []int, first int) []int {
	// make result array encoded's length + 1
	res := make([]int, len(encoded)+1)
	// fill first
	res[0] = first
	// go through encoded
	for i, n := range encoded {
		// original: encoded[i] = res[i] xor res[i+1]
		// derived: res[i+1] = encoded[i] xor res[i]
		res[i+1] = n ^ res[i]
	}

	return res
}
