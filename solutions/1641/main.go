package main

type Key struct {
	start int
	n     int
}

func countVowelStrings(n int) int {
	// cache for memoization
	cache := map[Key]int{}
	// start backtracking
	return backtrack(cache, 0, n)
}

// helper method to recursively backtrack
func backtrack(cache map[Key]int, start, n int) int {
	// check if we have a vowel string of length n
	if n == 0 {
		return 1
	}

	key := Key{start, n}

	// check if we have result cached already
	if v, ok := cache[key]; ok {
		return v
	}

	res := 0

	// simulate going through vowels a,e,i,o,u from start
	for i := start; i < 5; i++ {
		res += backtrack(cache, i, n-1)
	}

	cache[key] = res
	return res
}
