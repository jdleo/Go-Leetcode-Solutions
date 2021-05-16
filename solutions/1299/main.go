package main

func replaceElements(arr []int) []int {
	// current greatest
	max := -1

	// go thru array backwards
	for i := len(arr) - 1; i > -1; i-- {
		// set the max on the rightside to this num
		temp := arr[i]
		arr[i] = max

		if temp > max {
			max = temp
		}
	}

	return arr
}
