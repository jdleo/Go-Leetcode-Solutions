package main

func numTilePossibilities(tiles string) int {
	res := map[string]struct{}{}
	visited := make([]bool, len(tiles))
	// start backtracking
	backtrack(res, visited, tiles, "")
	return len(res) - 1
}

// helper method to recursively backtrack
func backtrack(res map[string]struct{}, visited []bool, tiles, path string) {
	// add this current path
	res[path] = struct{}{}
	// go thru possible tiles
	for i := 0; i < len(tiles); i++ {
		// check if we've already taken this char
		if visited[i] {
			continue
		}

		// visit this path
		visited[i] = true
		// backtrack
		backtrack(res, visited, tiles, path+string(tiles[i]))
		// un-visit this path
		visited[i] = false
	}
}
