package main

func shuffle(nums []int, n int) []int {
	res := append([]int{}, nums...)

	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[n+i]
	}

	return res
}
