package main

func canJump(nums []int) bool {
	// max index we can reach
	reach := 0

	// go through nums (up to last index)
	for i := 0; i < len(nums)-1; i++ {
		// update max reach we can jump to (if we can beat it)
		if i+nums[i] > reach {
			reach = i + nums[i]
		}

		// check if we are stuck here (currently on max reach, can't move)
		if i == reach {
			return false
		}
	}

	// we made it
	return true
}
