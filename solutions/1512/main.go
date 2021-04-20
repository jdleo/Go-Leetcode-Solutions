package main

func numIdenticalPairs(nums []int) int {
	counts := make([]int, 101)
	res := 0

	for _, num := range nums {
		res += counts[num]
		counts[num]++
	}

	return res
}
