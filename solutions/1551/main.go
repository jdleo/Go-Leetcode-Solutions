package main

func minOperations(n int) int {
	res := 0

	// go up to n, and by odds
	for i := 1; i < n; i += 2 {
		// we have to increase this number up to n
		res += n - i
	}

	return res
}
