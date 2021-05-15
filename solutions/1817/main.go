package main

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	res := make([]int, k)
	// users map of id -> map of minutes
	users := make(map[int]map[int]struct{})

	// go through logs
	for _, log := range logs {
		if users[log[0]] == nil {
			users[log[0]] = make(map[int]struct{})
		}

		// add this minute to user's map
		users[log[0]][log[1]] = struct{}{}
	}

	// go through users
	for _, uam := range users {
		// add length of unique minutes (uam) to res
		res[len(uam)-1]++
	}

	return res
}
