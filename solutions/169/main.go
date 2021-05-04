package main

func majorityElement(nums []int) int {
	// majority element, and count
	major, count := nums[0], 1

	// go through rest of nums
	for i := 1; i < len(nums); i++ {
		// check if current streak is 0
		if count == 0 {
			// increment and set majority
			count++
			major = nums[i]
		} else if major == nums[i] {
			// keep streak going
			count++
		} else {
			// decrease streak
			count--
		}
	}

	return major
}
