package main

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	// go until pointers meet
	for left <= right {
		// calculate mid point
		mid := (left + right) >> 1
		// check if left and right are both equal to target (found first n last)
		if nums[left] == nums[right] && nums[right] == target {
			return []int{left, right}
		} else if nums[mid] < target {
			// we need to search right half
			left = mid + 1
		} else if nums[mid] > target {
			// we need to search left half
			right = mid - 1
		} else {
			// mid IS target, close in on target
			if nums[left] != target {
				left++
			} else if nums[right] != target {
				right--
			}
		}
	}

	// never found
	return []int{-1, -1}
}
