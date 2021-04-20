package main

import "math"

func countPoints(points [][]int, queries [][]int) []int {
	res := make([]int, len(queries))

	for i, query := range queries {
		hits := 0
		x1, y1, r := query[0], query[1], float64(query[2])

		for _, point := range points {
			x2, y2 := point[0], point[1]

			if math.Sqrt(float64(((x1-x2)*(x1-x2))+((y1-y2)*(y1-y2)))) <= float64(r) {
				hits++
			}
		}
		res[i] = hits
	}

	return res
}
