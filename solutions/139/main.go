package main

type Key struct {
	path string
	cur  int
}

func wordBreak(s string, wordDict []string) bool {
	// convert word dict to set for o(1) lookups later
	set := map[string]bool{}
	for _, word := range wordDict {
		set[word] = true
	}

	// start backtracking
	cache := map[Key]bool{}
	return backtrack(cache, set, s, "", 0)
}

// helper method for recursively backtracking
func backtrack(cache map[Key]bool, set map[string]bool, s string, path string, cur int) bool {
	// check if we reached the end, and current path is cleared (all words accounted)
	if cur == len(s) {
		return path == ""
	}

	// check if already searched this path
	key := Key{path, cur}
	if v, ok := cache[key]; ok {
		if !v {
			return false
		}
	}

	// add this char to path
	path += string(s[cur])

	// check if we can break here
	if set[path] {
		// break by clearing path, backtrack
		if backtrack(cache, set, s, "", cur+1) {
			return true
		}
	}

	// also search path of not breaking here
	if !backtrack(cache, set, s, path, cur+1) {
		cache[key] = false
		return false
	} else {
		return true
	}
}
