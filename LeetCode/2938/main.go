package main

import (
	"fmt"
)

func minimumSteps(s string) int64 {
	res, z := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		if char == '0' {
			z++
		} else {
			res += z
		}
	}
	return int64(res)
}

func main() {
	var s string
	fmt.Println("Enter a binary string:")
	fmt.Scanln(&s)
	steps := minimumSteps(s)
	fmt.Printf("Minimum steps required: %d\n", steps)
}
