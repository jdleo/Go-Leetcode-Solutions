package main

func sumZero(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i*2 - n + 1
	}
	return res
}
