package main

func runningSum(nums []int) []int {
	res := nums[:]
	for i := 1; i < len(nums); i++ {
		res[i] += res[i-1]
	}
	return res
}
