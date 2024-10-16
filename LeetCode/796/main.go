package main

import (
	"fmt"
)

func rotateString(s string, goal string) bool {
	for i := 0; i < len(s); i++ {
		s = shift(s)
		if s == goal {
			return true
		}
	}
	return false
}

func shift(s string) string {
	r := ""
	for i := 1; i < len(s); i++ {
		r = r + string(s[i])
	}
	return r + string(s[0])
}

func main() {
	s := "abcde"
	goal := "cdeab"
	if rotateString(s, goal) {
		fmt.Printf("Yes, '%s' is a rotation of '%s'.\n", goal, s)
	} else {
		fmt.Printf("No, '%s' is not a rotation of '%s'.\n", goal, s)
	}
}
