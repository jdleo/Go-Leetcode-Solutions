package main

func numJewelsInStones(jewels string, stones string) int {
	set := map[rune]bool{}
	res := 0
	for _, jewel := range jewels {
		set[jewel] = true
	}

	for _, stone := range stones {
		if _, ok := set[stone]; ok {
			res++
		}
	}

	return res
}
