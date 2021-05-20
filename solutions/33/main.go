package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	// go until pointers meet
	for left <= right {
		// get mid point
		mid := (left + right) >> 1

		// check if we found
		if nums[mid] == target {
			return mid
		}

		// check if left side is sorted
		if nums[left] <= nums[mid] {
			// check if target is in left side
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// check if target is in right side
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
