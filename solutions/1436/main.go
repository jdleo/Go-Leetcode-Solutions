package main

func destCity(paths [][]string) string {
	// map to keep track of in/out degrees
	inout := make(map[string][]int)

	// go through paths
	for _, path := range paths {
		if inout[path[0]] == nil {
			inout[path[0]] = []int{0, 0}
		}
		if inout[path[1]] == nil {
			inout[path[1]] = []int{0, 0}
		}

		// increment in/out degree
		inout[path[0]][1]++
		inout[path[1]][0]++
	}

	// go thru in/out degrees
	for city, degree := range inout {
		// the one with outdegree of 0 is the destination
		if degree[1] == 0 {
			return city
		}
	}

	return ""
}
