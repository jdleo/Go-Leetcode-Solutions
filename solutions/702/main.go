package main

// This is the ArrayReader's API interface.
// You should not implement it, or speculate about its implementation
type ArrayReader struct {
}

func (this *ArrayReader) get(index int) int { return 0 }

func search(reader ArrayReader, target int) int {
	right := 1
	// keep left shifting right, until its bigger than target
	for reader.get(right) < target {
		right <<= 1
	}
	// make left pointer the right, right shifted by 1 (smaller)
	left := right >> 1

	// now we can binary search
	for left <= right {
		// get mid point
		mid := (left + right) >> 1

		// check if this is our target
		n := reader.get(mid)
		if n == target {
			return mid
		} else if n < target {
			// search right
			left = mid + 1
		} else {
			// search left
			right = mid - 1
		}
	}

	return -1
}
