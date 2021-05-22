package main

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	// sort by heights descending, queue position ascending
	sort.Slice(people, func(i, j int) bool {
		// if equal heights, sort queue position
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		// sort by height
		return people[i][0] > people[j][0]
	})

	res := make([][]int, len(people))

	// go thru sorted people
	for _, person := range people {
		// get index to insert at, and insert at index
		index := person[1]
		res = append(res[:index+1], res[index:]...)
		res[index] = person
	}

	return res[:len(people)]
}
