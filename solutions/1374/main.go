package main

func generateTheString(n int) string {
	res := make([]byte, n)
	// check if even (need to add 1 'y')
	if n&1 == 0 {
		res[0] = 'y'
	}

	// start from 1 if it was even, fill rest of characters
	for i := (n & 1) ^ 1; i < n; i++ {
		res[i] = 'x'
	}

	return string(res)
}
