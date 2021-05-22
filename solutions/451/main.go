package main

import "sort"

func frequencySort(s string) string {
	counts := make([]int, 255)
	res := []byte(s)

	// count char freq in s
	for _, char := range s {
		counts[char-'0']++
	}

	// sort based on custom comparator
	sort.Slice(res, func(i, j int) bool {
		// if equal freq, sort based on char
		if counts[res[i]-'0'] == counts[res[j]-'0'] {
			return res[i] < res[j]
		}
		return counts[res[i]-'0'] > counts[res[j]-'0']
	})

	return string(res)
}
