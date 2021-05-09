package main

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, n)
	base3, base5, base15 := 3, 5, 15

	for i := 1; i <= n; i++ {
		if i == base15 {
			res[i-1] = "FizzBuzz"
			base3, base5, base15 = base3+3, base5+5, base15+15
		} else if i == base3 {
			res[i-1] = "Fizz"
			base3 += 3
		} else if i == base5 {
			res[i-1] = "Buzz"
			base5 += 5
		} else {
			res[i-1] = strconv.Itoa(i)
		}
	}

	return res
}
