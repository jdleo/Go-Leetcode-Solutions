package main

func numIdenticalPairs(nums []int) int {
	hm := map[int]int{}
	res := 0

	for _, num := range nums {
		if val, ok := hm[num]; ok {
			res += val
			hm[num]++
		} else {
			hm[num] = 1
		}
	}

	return res
}
