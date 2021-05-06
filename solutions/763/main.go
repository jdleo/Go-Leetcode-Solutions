package main

func partitionLabels(S string) []int {
	res := []int{}

	// array to keep track of last index
	lastIndex := make([]int, 26)

	// mark the last index that each char is seen at
	for i, char := range S {
		lastIndex[char-'a'] = i
	}

	// start and end
	start, end := 0, 0

	// go through string
	for i, char := range S {
		// update max end we can reach
		if lastIndex[char-'a'] > end {
			end = lastIndex[char-'a']
		}

		// check if we reached max end (this means we have all characters we need)
		if i == end {
			// this is a valid substring
			res = append(res, end-start+1)
			start = i + 1
		}
	}

	return res
}
