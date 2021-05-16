package main

func finalPrices(prices []int) []int {
	stack := []int{}

	// go through prices
	for i, price := range prices {
		// check if we can use current price AS a discount for SOMETHING
		for len(stack) > 0 && prices[stack[len(stack)-1]] >= price {
			// apply discount, and pop from stack
			prices[stack[len(stack)-1]] -= price
			stack = stack[:len(stack)-1]
		}
		// add index to stack
		stack = append(stack, i)
	}

	// discounted prices
	return prices
}
