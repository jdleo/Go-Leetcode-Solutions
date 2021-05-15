package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	res := 0

	// go through array
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// check first condition
			if abs(arr[i]-arr[j]) > a {
				// we dont have to continue
				continue
			}
			for k := j + 1; k < n; k++ {
				// check other two conditions
				if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					res++
				}
			}
		}
	}

	return res
}

// helper method for abs
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
