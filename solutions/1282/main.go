package main

func groupThePeople(groupSizes []int) [][]int {
	res := [][]int{}
	// individual group buckets
	groups := make(map[int][]int)

	// go through people
	for i := 0; i < len(groupSizes); i++ {
		// add this person to the bucket
		groups[groupSizes[i]] = append(groups[groupSizes[i]], i)
	}

	// go through group buckets
	for size, group := range groups {
		// temp array
		temp := []int{}

		// go through group
		for _, person := range group {
			// add person to temp
			temp = append(temp, person)

			// check if we're at max group capacity
			if len(temp) == size {
				// add to res and clear temp
				res = append(res, temp)
				temp = []int{}
			}
		}
	}

	return res
}
