package main

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	res := 0

	// go until num is 0
	for num > 0 {
		// divide by 2 no matter what, (thats a step)
		// add extra 1 if this is odd
		res += 1 + num&1
		// right shift num
		num >>= 1
	}

	// we do -1 because the last step would have been 1->0 and it was counted as 2
	return res - 1
}
