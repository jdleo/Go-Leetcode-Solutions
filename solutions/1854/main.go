package main

func maximumPopulation(logs [][]int) int {
	// maps to count births and deaths
	births, deaths := map[int]int{}, map[int]int{}

	// go through logs
	for _, log := range logs {
		// increment birth for this year, increment death for this year
		births[log[0]]++
		deaths[log[1]]++
	}

	// max population, current population, and year
	maxPopulation, population, year := 0, 0, 0

	// go through possible years
	for i := 1950; i < 2051; i++ {
		// increase/decrease population
		population += births[i] - deaths[i]
		// check if population is bigger than max
		if population > maxPopulation {
			// set max population and year
			maxPopulation, year = population, i
		}
	}

	return year
}
