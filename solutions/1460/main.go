package main

func canBeEqual(target []int, arr []int) bool {
	// frequency of each num in target
	counts := map[int]int{}

	// go thru target, and increment
	for _, num := range target {
		counts[num]++
	}

	// go thru arr, and decrement
	for _, num := range arr {
		// if after decrementing, check if its less than 0
		if counts[num]--; counts[num] < 0 {
			// can't make
			return false
		}
	}

	return true
}
