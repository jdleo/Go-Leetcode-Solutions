package main

func sumOddLengthSubarrays(arr []int) int {
	res := 0

	// go through all possible odd lengths
	for i := 1; i <= len(arr); i += 2 {
		// go through all possible subarrays
		for j := 0; j < len(arr)-i+1; j++ {
			// add sum of odd-length subarray
			res += sum(arr[j : j+i])
		}
	}

	return res
}

// helper method to get sum of slice
func sum(arr []int) int {
	res := 0

	for _, num := range arr {
		res += num
	}

	return res
}
