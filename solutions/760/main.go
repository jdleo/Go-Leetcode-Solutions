package main

func anagramMappings(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums2))
	// mapping of nums to indices for nums2
	seen := make(map[int]int, len(nums2))

	// fill map
	for i, num := range nums2 {
		seen[num] = i
	}

	// go thru nums1
	for i, num := range nums1 {
		// get where it's seen in nums2
		res[i] = seen[num]
	}

	return res
}
