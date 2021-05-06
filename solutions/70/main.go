package main

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	// dinstinct ways 1 step ago, 2 steps ago
	a, b := 1, 1

	// calculate rest
	for i := 2; i <= n; i++ {
		a, b = a+b, a
	}

	return a
}
