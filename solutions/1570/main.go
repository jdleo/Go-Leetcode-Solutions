package main

type SparseVector struct {
	hm map[int]int
}

func Constructor(nums []int) SparseVector {
	// hash map to hold nonzero indices -> num
	hm := make(map[int]int, len(nums))

	for i, num := range nums {
		if num != 0 {
			hm[i] = num
		}
	}

	return SparseVector{hm}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	res := 0
	// go through all nonzero indices in this map
	for i, num := range this.hm {
		// check if index also exists in other vector
		if vec.hm[i] != 0 {
			res += num * vec.hm[i]
		}
	}

	return res
}
