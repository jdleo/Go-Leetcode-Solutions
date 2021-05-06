package main

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	// map for digits to letters
	letters := []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	// start backtracking
	backtrack(&res, &[]byte{}, letters, digits, 0)

	return res
}

func backtrack(res *[]string, path *[]byte, letters []string, digits string, cur int) {
	// base case (found a combination)
	if len(*path) == len(digits) {
		*res = append(*res, string(*path))
		return
	}

	// go through all possible letters, in current digit
	candidates := letters[digits[cur]-'0']
	for i := 0; i < len(candidates); i++ {
		// take this character
		*path = append(*path, candidates[i])
		// backtrack
		backtrack(res, path, letters, digits, cur+1)
		// remove this character
		*path = (*path)[:len(*path)-1]
	}

}
