package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	max := 0

	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}

	for i := 0; i < len(candies); i++ {
		res[i] = candies[i]+extraCandies >= max
	}

	return res
}
