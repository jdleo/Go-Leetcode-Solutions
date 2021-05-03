package main

func createTargetArray(nums []int, index []int) []int {
	res := []int{}

	// go through index
	for k, v := range index {
		if len(res) == index[k] {
			res = append(res, nums[index[k]])
		} else {
			res = append(res, 0)
			copy(res[v+1:], res[v:])
			res[index[k]] = nums[k]

		}
	}
	return res
}
