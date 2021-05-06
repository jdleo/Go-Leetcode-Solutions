package main

func numTeams(rating []int) int {
	res := 0

	// go through ratings
	for j := 1; j < len(rating)-1; j++ {
		left_smaller, left_bigger := 0, 0
		right_smaller, right_bigger := 0, 0

		// check those to the LEFT of j (the i's)
		for i := 0; i < j; i++ {
			if rating[i] < rating[j] {
				left_smaller++
			} else {
				left_bigger++
			}
		}

		// check those to the RIGHT of j (the k's)
		for k := j + 1; k < len(rating); k++ {
			if rating[k] < rating[j] {
				right_smaller++
			} else {
				right_bigger++
			}
		}

		// result is just the combination of teams that can be formed
		res += left_smaller*right_bigger + left_bigger*right_smaller

	}

	return res
}
