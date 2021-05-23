package main

func fib(n int) int {
	if n <= 1 {
		return n
	}

	// this represents fib(n-1) and fib(n-2)
	a, b := 1, 0

	// calculate 2 to n
	for i := 2; i <= n; i++ {
		// fib(n-1) + fib(n-2)
		a, b = a+b, a
	}

	return a
}
