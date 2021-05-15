package main

func largestAltitude(gain []int) int {
	alt, res := 0, 0

	// go through gains
	for _, g := range gain {
		// ascend/descend
		alt += g
		// check if max
		if alt > res {
			res = alt
		}
	}

	return res
}
