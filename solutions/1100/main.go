package main

func numKLenSubstrNoRepeats(S string, K int) int {
	res := 0
	counts := map[byte]int{}

	// go through string S
	for i := 0; i < len(S); i++ {
		// increment count for this character
		counts[S[i]]++

		// check if index is past K length
		if i >= K {
			// get the caracter we have to remove (sliding window of size k)
			removed := S[i-K]
			// decrement in counts
			counts[removed]--
			// check if no more of that character left (just delete from map)
			if counts[removed] == 0 {
				delete(counts, removed)
			}
		}

		// check if we have valid substring
		if len(counts) == K {
			res++
		}
	}

	return res
}
