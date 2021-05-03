package main

import "strings"

func interpret(command string) string {
	/**
	() -> o
	(al) -> al
	*/
	return strings.ReplaceAll(strings.ReplaceAll(command, "(al)", "al"), "()", "o")
}
