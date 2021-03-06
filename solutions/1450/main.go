package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	res := 0

	for i := 0; i < len(startTime); i++ {
		// check if query time lies within this student's interval
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			res++
		}
	}

	return res
}
