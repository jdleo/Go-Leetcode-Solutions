package main

func findAnagrams(s string, p string) []int {
	res, need, have := []int{}, [26]int{}, [26]int{}
	start := 0

	// check if s is smaller than p (impossible)
	if len(s) < len(p) {
		return res
	}

	// fill need and have
	for i := 0; i < len(p); i++ {
		need[p[i]-'a']++
		have[s[i]-'a']++
	}

	// start at length of p, go til end of s
	for i := len(p); i < len(s); i, start = i+1, start+1 {
		// check if anagrams
		if need == have {
			res = append(res, start)
		}

		// increment/decrement have
		have[s[i]-'a']++
		have[s[start]-'a']--
	}

	// check if anagrams
	if need == have {
		res = append(res, start)
	}

	// start at length
	return res
}
