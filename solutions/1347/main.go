package main

func minSteps(s string, t string) int {
	// counts for characters
	counts := make([]int, 26)
	// fill counts from both strings
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	res := 0

	// go through counts
	for i := 0; i < 26; i++ {
		// count positives
		if counts[i] > 0 {
			res += counts[i]
		}
	}

	return res
}
