package main

func selfDividingNumbers(left int, right int) []int {
	res := []int{}

	// left to right
	for i := left; i < right+1; i++ {
		if isSelfDividing(i) {
			res = append(res, i)
		}
	}

	return res
}

func isSelfDividing(n int) bool {
	x := n
	for n > 0 {
		if n%10 == 0 || x%(n%10) != 0 {
			return false
		}
		n /= 10
	}
	return true
}
