package main

import "math"

func romanToInt(s string) int {
	res, last := 0, math.MaxInt64

	// roman mappings
	romans := map[rune]int{
		'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100, 'D': 500,
		'M': 1_000,
	}

	for _, char := range s {
		// check if last char value was smaller than cur (ex: IV = 4)
		if last < romans[char] {
			// we need to subtract 2*last
			res -= last * 2
		}
		// add this char's roman value
		res += romans[char]
		// set last
		last = romans[char]
	}

	return res
}
