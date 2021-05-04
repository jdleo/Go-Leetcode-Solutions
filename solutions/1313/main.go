package main

func decompressRLElist(nums []int) []int {
	res := []int{}

	// loop through nums by 2
	for i := 0; i < len(nums); i += 2 {
		freq, val := nums[i], nums[i+1]
		// fill value 'val', 'freq' times
		for freq > 0 {
			res = append(res, val)
			freq--
		}
	}

	return res
}
