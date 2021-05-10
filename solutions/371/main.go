package main

func getSum(a int, b int) int {
	// if b is 0, its just a
	if b == 0 {
		return a
	}

	// get the carry bit
	carry := (a & b) << 1
	// sum
	sum := a ^ b
	// pass sum and carry
	return getSum(sum, carry)
}
