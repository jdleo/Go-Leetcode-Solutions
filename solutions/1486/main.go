package main

func xorOperation(n int, start int) int {
	res := 0

	// self explanatory :D
	for i := start; i < start+n*2; i += 2 {
		res ^= i
	}

	return res
}
