package main

func minCostToMoveChips(position []int) int {
	// counts for even/odd position chips
	parity := []int{0, 0}

	// count parities
	for _, p := range position {
		parity[p%2]++
	}

	// return minimum moves of both
	if parity[0] < parity[1] {
		return parity[0]
	}
	return parity[1]
}
