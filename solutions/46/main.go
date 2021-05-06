package main

func permute(nums []int) [][]int {
	res := [][]int{}
	// start backtracking, pass reference to res
	backtrack(&res, nums, []int{}, map[int]bool{})
	return res
}

// helper method to generate permutations
func backtrack(res *[][]int, nums []int, temp []int, used map[int]bool) {
	// check if we have a valid perm
	if len(temp) == len(nums) {
		// add copy of temp to res
		tmp := make([]int, len(nums))
		copy(tmp, temp)
		*res = append(*res, tmp[:])
		return
	}

	// go through all possible nums
	for _, num := range nums {
		// check if we didnt already use this num
		if !used[num] {
			// add to used/temp, backtrack
			used[num] = true
			temp = append(temp, num)
			backtrack(res, nums, temp, used)
			// remove from used/temp
			delete(used, num)
			temp = temp[:len(temp)-1]
		}
	}
}
