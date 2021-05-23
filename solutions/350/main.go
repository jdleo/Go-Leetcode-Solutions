package main

func intersect(nums1 []int, nums2 []int) []int {
	// counts of frequencies
	counts := map[int]int{}

	// fill counts
	for _, num := range nums1 {
		counts[num]++
	}

	res := []int{}

	// go thru nums2
	for _, num := range nums2 {
		// check if still numbers to take
		if counts[num] > 0 {
			// take, decrement count
			res = append(res, num)
			counts[num]--
		}
	}

	return res
}
