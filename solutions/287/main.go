package main

func findDuplicate(nums []int) int {
	// tortoise and hare
	slow := nums[0]
	fast := nums[nums[0]]

	// go til they meet
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// reset fast
	fast = 0
	// go til they meet
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
