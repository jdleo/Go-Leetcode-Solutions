package main

func subsetXORSum(nums []int) int {
	var res int
	backtrack(&res, nums, []int{}, 0)
	return res
}

func backtrack(res *int, nums []int, path []int, start int) {
	var temp int
	// xor all numbers in current subset to temp
	for _, num := range path {
		temp ^= num
	}
	// add to res
	*res += temp

	// go through all possible candidates
	for i := start; i < len(nums); i++ {
		// backtrack
		path = append(path, nums[i])
		backtrack(res, nums, path, i+1)
		path = path[:len(path)-1]
	}
}
