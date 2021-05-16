package main

func sumBase(n int, k int) int {
	res := 0
	// go thru each digit in base k representation and sum
	for _, num := range baseK(n, k) {
		res += num
	}
	return res
}

// helper method to get base k of int n as int slice
// (reversed order, but that doesn't matter)
func baseK(n, k int) []int {
	res := []int{}

	// go until n is 0
	for n > 0 {
		// add the current num mod k
		res = append(res, n%k)
		// divide by k
		n /= k
	}

	return res
}
