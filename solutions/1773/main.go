package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	// map each rule key to its index in items array
	ruleKeys := map[string]int{"type": 0, "color": 1, "name": 2}
	// go through items
	for _, item := range items {
		// check if the thing at current rule key matches rule value
		if item[ruleKeys[ruleKey]] == ruleValue {
			res++
		}
	}
	return res
}
