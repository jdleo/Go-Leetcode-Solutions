package main

func maxArea(height []int) int {
	res, left, right := 0, 0, len(height)-1

	for left < right {
		// check the lower wall
		if height[left] < height[right] {
			// set new max area, and move left forward
			if height[left]*(right-left) > res {
				res = height[left] * (right - left)
			}
			left++
		} else {
			// set new max area, and move right forward
			if height[right]*(right-left) > res {
				res = height[right] * (right - left)
			}
			right--
		}
	}

	return res
}
