package main

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, days[len(days)-1]+1)
	set := make(map[int]bool)

	// fill set for o(1) lookups later
	for _, day := range days {
		set[day] = true
	}

	// go through all possible days
	for day := 1; day < days[len(days)-1]+1; day++ {
		// check if not in our set
		if !set[day] {
			// just take last day
			dp[day] = dp[day-1]
		} else {
			// calculate costs of all three tickets, + going 1,7,30 days back
			day1 := costs[0] + dp[max(day-1, 0)]
			day7 := costs[1] + dp[max(day-7, 0)]
			day30 := costs[2] + dp[max(day-30, 0)]

			// min cost of all of these
			dp[day] = min(day1, min(day7, day30))
		}
	}

	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
