package main

import "math"

func maxProfit(prices []int) int {
	res, min := 0, math.MaxInt64

	// go thru prices
	for _, price := range prices {
		// check if we can take a profit
		if price > min {
			// take profit, and set new min
			res += price - min
			min = price
		}

		// otherwise, just set new min
		if price < min {
			min = price
		}
	}

	return res
}
