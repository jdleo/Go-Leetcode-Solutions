package main

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	// sort nums
	sort.Ints(nums)
	// go through numbers for i
	for i := 0; i < len(nums); i++ {
		// check if duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// get j and k pointers
		j, k := i+1, len(nums)-1

		// go until pointers meet
		for j < k {
			// get sum
			sum := nums[i] + nums[j] + nums[k]
			// if less than 0, we need to move lower pointer up
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				// this is a valid i+j+k=0
				res = append(res, []int{nums[i], nums[j], nums[k]})
				// to avoid duplicates, move j and k forward
				for j+1 > k && nums[j] == nums[j+1] {
					j++
				}
				for k-1 > j && nums[k] == nums[k-1] {
					k--
				}
				j, k = j+1, k-1
			}
		}
	}

	return res
}
