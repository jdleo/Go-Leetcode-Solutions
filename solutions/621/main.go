package main

import "sort"

func leastInterval(tasks []byte, n int) int {
	// base case
	if len(tasks) == 0 {
		return 0
	}

	// task frequencies
	freq := make([]int, 26)

	// count tasks
	for i := 0; i < len(tasks); i++ {
		freq[tasks[i]-'A']++
	}

	// sort tasks low->high, and also get max frequent task (all that matters)
	sort.Ints(freq)
	max := freq[25] - 1
	idle := max * n // min amount of time that max freq will cause idle

	// go thru rest of tasks
	for i := 0; i < 25; i++ {
		// reduce idle time (these tasks fill the idle space)
		idle -= min(max, freq[i])
	}

	// edge case : idle was negative
	if idle < 0 {
		idle = 0
	}

	// total wait time is tasks length + remaining idle time not filled
	// and -1 because one of the tasks doesnt have to wait
	return len(tasks) + idle

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
