package main

import (
	"fmt"
)

func minAddToMakeValid(s string) int {
	var stack []int
	var c int
	for _, r := range s {
		if r == '(' {
			stack = append(stack, 1)
		} else if r == ')' {
			if len(stack) == 0 {
				c++
				continue
			}
			stack = stack[1:]
		}
	}
	c += len(stack)
	return c
}

func main() {
	s1 := "())"
	s2 := "((("
	s3 := "()"
	s4 := "()))(("

	fmt.Println("Input:", s1, " Min Additions:", minAddToMakeValid(s1)) // Output: 1
	fmt.Println("Input:", s2, " Min Additions:", minAddToMakeValid(s2)) // Output: 3
	fmt.Println("Input:", s3, " Min Additions:", minAddToMakeValid(s3)) // Output: 0
	fmt.Println("Input:", s4, " Min Additions:", minAddToMakeValid(s4))  // Output: 4
}
