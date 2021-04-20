package main

func maximumWealth(accounts [][]int) int {
	res := 0

	for _, account := range accounts {
		sum := 0
		for _, bal := range account {
			sum += bal
		}

		if sum > res {
			res = sum
		}
	}
	return res
}
