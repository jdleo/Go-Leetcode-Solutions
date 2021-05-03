package main

func subtractProductAndSum(n int) int {
	// product and sum
	product, sum := 1, 0

	// go until no more digits
	for n > 0 {
		// multiply by rightmost digit, add rightmost digit
		product *= n % 10
		sum += n % 10
		// shift to get next digit
		n /= 10
	}

	return product - sum
}
