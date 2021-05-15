package main

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		for i := 0; i <= (len(row)-1)/2; i++ {
			row[i], row[len(row)-1-i] = row[len(row)-1-i]^1, row[i]^1
		}
	}
	return A
}
