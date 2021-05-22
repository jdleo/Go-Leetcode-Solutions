package main

import (
	"sort"
)

func customSortString(order string, str string) string {
	// make str a byte slice
	res := []byte(str)

	// fill characters of order, based on where they were seen in order
	index := make([]int, 26)
	for i, char := range order {
		index[char-'a'] = i
	}

	// sort result string based on index mapping
	sort.Slice(res, func(i, j int) bool {
		return index[res[i]-'a'] < index[res[j]-'a']
	})

	return string(res)
}
