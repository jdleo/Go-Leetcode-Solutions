package main

func sortString(s string) string {
	counts := make([]int, 26)
	res := []byte{}

	// fill counts
	for _, char := range s {
		counts[char-'a']++
	}

	// return byte slice as string
	return string(res)
}
