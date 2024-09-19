package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	var result []int
	for i := 0; i < len(expression); i++ {
		char := expression[i]
		if char == '+' || char == '-' || char == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch char {
					case '+':
						result = append(result, l+r)
					case '-':
						result = append(result, l-r)
					case '*':
						result = append(result, l*r)
					}
				}
			}
		}
	}

	if len(result) == 0 {
		num, _ := strconv.Atoi(expression)
		result = append(result, num)
	}

	return result
}

func main() {
	expression := "2*3-4*5"
	result := diffWaysToCompute(expression)
	fmt.Println("Possible results for the expression:", result)
}
