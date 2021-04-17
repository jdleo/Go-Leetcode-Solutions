package main

import "strings"

func defangIPaddr(address string) string {
	var sb strings.Builder

	for _, c := range address {
		if c == '.' {
			sb.WriteString("[.]")
		} else {
			sb.WriteRune(c)
		}
	}

	return sb.String()
}
