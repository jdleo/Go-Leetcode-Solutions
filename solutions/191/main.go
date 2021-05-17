package main

func hammingWeight(num uint32) int {
	var res uint32 = 0
	// go until no more bits to check
	for num > 0 {
		// this will be 1 if rightmost bit is 1
		res += num & 1
		num >>= 1
	}
	return int(res)
}
