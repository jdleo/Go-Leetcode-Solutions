package main

func moveZeroes(nums []int) {
	// zero index
	zero := 0

	// go through nums
	for i := 0; i < len(nums); i++ {
		// check if current num is nonzero
		if nums[i] != 0 {
			// swap with the zero index
			nums[zero], nums[i] = nums[i], nums[zero]
			// move last zero index seen forward
			zero++
		}
	}
}
