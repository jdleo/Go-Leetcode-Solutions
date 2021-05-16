package main

func sumOfUnique(nums []int) int {
	counts, res := map[int]int{}, 0

	// go thru nums, and count individuals
	for _, num := range nums {
		counts[num]++
	}

	// go thru counts, and count those unique
	for _, count := range counts {
		if count == 1 {
			res++
		}
	}

	return res
}
