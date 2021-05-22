package main

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	seen := map[int]bool{}

	// mark nums as seen
	for _, num := range nums1 {
		seen[num] = true
	}

	// go thru nums2
	for _, num := range nums2 {
		// check if we've seen this number
		if seen[num] {
			// add to res, and mark as not seen to ensure uniqueness
			res = append(res, num)
			seen[num] = false
		}
	}

	return res
}
