package main

func sumOfDigits(nums []int) int {
	min := nums[0]

	// find the minimum number
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	sum := 0

	// get sum of digits
	for min > 0 {
		sum += min % 10
		min /= 10
	}

	// 0 if odd, 1 if even
	return 1 ^ (sum & 1)
}
