package main

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	// go until pointers meet
	for left < right {
		// get chars
		a, b := s[left], s[right]
		// check if left is punctuation
		if isPunctuation(a) {
			left++
			continue
		}
		if isPunctuation(b) {
			right--
			continue
		}

		// move to lowercase
		lower(&a)
		lower(&b)

		// make sure equal
		if a != b {
			return false
		}

		left, right = left+1, right-1
	}

	// no issues
	return true
}

func lower(a *byte) {
	// check if uppercase
	if *a >= 'A' && *a <= 'Z' {
		*a += 32
	}
}

func isPunctuation(a byte) bool {
	return !(('A' <= a && a <= 'Z') || ('z' <= a && a <= 'z') || ('0' <= a && a <= '9'))
}
