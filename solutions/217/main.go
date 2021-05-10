package main

func containsDuplicate(nums []int) bool {
	// map to keep track of seen nums
	seen := map[int]bool{}

	for _, num := range nums {
		// check if we've seen this num before
		if seen[num] {
			return true
		}

		seen[num] = true
	}

	return false
}
