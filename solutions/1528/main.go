package main

func restoreString(s string, indices []int) string {
	// result byte slice
	res := make([]byte, len(s))

	// go through string
	for i := 0; i < len(s); i++ {
		// insert byte i, at indices i
		res[indices[i]] = s[i]
	}

	// result as string
	return string(res)
}
