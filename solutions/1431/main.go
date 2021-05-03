package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := []bool{}
	max := 0

	for _, c := range candies {
		if c > max {
			max = c
		}
	}

	for _, c := range candies {
		res = append(res, c+extraCandies >= max)
	}

	return res
}
