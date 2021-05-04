package main

func shuffle(nums []int, n int) []int {

	for i, j := n-1, len(nums)-1; j >= n; j, i = j-1, i-1 {
		nums[j] <<= 10
		nums[j] |= nums[i]
	}

	for i, j := 0, n; j < len(nums); i, j = i+2, j+1 {
		x := nums[j] & 0b1111111111
		y := nums[j] >> 10
		nums[i], nums[i+1] = x, y
	}

	return nums
}
