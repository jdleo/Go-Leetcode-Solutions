package main

func removeOuterParentheses(S string) string {
	res := []rune{}
	bal := 0

	for _, char := range S {
		if char == '(' {
			if bal > 0 {
				res = append(res, '(')
			}
			bal++
		} else {
			if bal > 1 {
				res = append(res, ')')
			}
			bal--
		}
	}

	return string(res)
}
