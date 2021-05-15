package main

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	res := []int{}
	i, j, k := 0, 0, 0

	// go through arrays simultaneously
	for i < len(arr1) && j < len(arr2) && k < len(arr3) {
		// get minimum
		curMin := min(arr1[i], min(arr2[j], arr3[k]))

		// check if all are equal
		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			res = append(res, arr1[i])
		}

		// increment minimum numbers

		if arr1[i] == curMin {
			i++
		}

		if arr2[j] == curMin {
			j++
		}

		if arr3[k] == curMin {
			k++
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
