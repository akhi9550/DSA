package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	examples := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	for _, example := range examples {
		if isValid(example) {
			fmt.Println(example, "is valid.")
		} else {
			fmt.Println(example, "is not valid.")
		}
	}
}
