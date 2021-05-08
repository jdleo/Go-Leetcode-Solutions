package main

func twoSum(nums []int, target int) []int {
	// map of numbers we have seen, with their index
	seen := map[int]int{}

	// go thru nums
	for i, num := range nums {
		// num + x = target
		// x = target - num
		if val, ok := seen[target-num]; ok {
			// found match
			return []int{val, i}
		}
		// write to seen
		seen[num] = i
	}

	// never found :o
	return []int{-1, -1}
}
