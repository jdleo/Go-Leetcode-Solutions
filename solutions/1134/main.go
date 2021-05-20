package main

import (
	"math"
)

func isArmstrong(n int) bool {
	k := math.Floor(math.Log10(float64(n)) + 1)
	res, x := 0.0, n

	for x > 0 {
		// add rightmost digit to the power of k
		res += math.Pow(float64(x%10), k)
		// move digit over
		x /= 10
	}

	// check if sum of powers is equal to n
	return res == float64(n)
}
