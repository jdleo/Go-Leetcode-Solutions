package main

func arraySign(nums []int) int {
	negatives := 0

	for _, num := range nums {
		// check if 0 (can early terminate)
		if num == 0 {
			return 0
		} else if num < 0 {
			// count neg
			negatives++
		}
	}

	// if even negatives - positive, if odd - negative
	if negatives&1 == 0 {
		return 1
	} else {
		return -1
	}
}
