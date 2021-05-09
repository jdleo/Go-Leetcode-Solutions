package main

func reverseString(s []byte) {
	n := len(s)
	// go up until midpoint
	for i := 0; i < n/2; i++ {
		// swap i from front, with i from back
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}
