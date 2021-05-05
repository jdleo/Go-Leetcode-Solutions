package main

func maxProfit(prices []int) int {
	res, curMin := 0, int(^uint(0)>>1)

	// go thru prices
	for _, price := range prices {
		// see if this is a greater profit than current max profit
		if price-curMin > res {
			res = price - curMin
		}

		// set new current minimum
		if price < curMin {
			curMin = price
		}
	}

	return res
}
