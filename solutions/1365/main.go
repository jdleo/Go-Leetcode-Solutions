package main

func smallerNumbersThanCurrent(nums []int) []int {
	// counts array (numbers are bounded in range [0,101])
	counts := make([]int, 101)
	// count each number's occurrence
	for _, num := range nums {
		counts[num]++
	}

	// running sum from left to right
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	// go through nums for answer
	for i := 0; i < len(nums); i++ {
		// if 0, nothing is smaller
		if nums[i] != 0 {
			// counts to the left of this number is total sum of counts of numbers smaller than current
			nums[i] = counts[nums[i]-1]
		}
	}

	return nums
}
