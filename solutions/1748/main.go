package main

func sumOfUnique(nums []int) int {
	counts, res := map[int]int{}, 0

	// go thru nums, and count individuals
	for _, num := range nums {
		counts[num]++
	}

	// go thru counts, and sum the unique ones
	for num, count := range counts {
		if count == 1 {
			res += num
		}
	}

	return res
}
