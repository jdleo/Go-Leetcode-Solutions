package main

func generateParenthesis(n int) []string {
	res := []string{}
	// start recursion
	helper(&res, "", n, n)
	return res
}

// helper method to generate paranthesis recursively
func helper(res *[]string, path string, left int, right int) {
	// check if we have a valid paranthesis string
	if left == 0 && right == 0 {
		*res = append(*res, path)
	} else {
		// check if we can take from left
		if left > 0 {
			helper(res, path+"(", left-1, right)
		}

		// check if we can take from right
		if right > left {
			helper(res, path+")", left, right-1)
		}
	}
}
