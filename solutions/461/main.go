package main

func hammingDistance(x int, y int) int {
	var res int

	// go until we still have bits to compare
	for x != 0 || y != 0 {
		// xor rightmost bits
		res += (x & 1) ^ (y & 1)
		// right shift both nums
		x, y = x>>1, y>>1
	}

	return res
}
