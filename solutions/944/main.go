package main

func minDeletionSize(strs []string) int {
	res := 0

	for i := 0; i < len(strs[0]); i++ {
		var last byte = 0
		// go through column
		for j := 0; j < len(strs); j++ {
			// make sure current byte is not smaller than last
			if strs[j][i] < last {
				res++
				break
			}
			last = strs[j][i]
		}
	}

	return res
}
