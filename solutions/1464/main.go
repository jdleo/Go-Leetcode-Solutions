package main

func maxProduct(nums []int) int {
	// first and second biggest
	first, second := -1, -1

	for _, num := range nums {
		// set first, second biggest
		if num > first {
			first, second = num, first
		} else if num > second {
			second = num
		}
	}

	// first-1 * second-1
	return (first - 1) * (second - 1)
}
