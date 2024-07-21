package main

import (
	"fmt"
)

func climbStairs(n int) int {
	a, b := 1, 1
	for ; n > 1; n-- {
		a, b = b, a+b
	}
	return b
}

func main() {
	steps := 5
	result := climbStairs(steps)
	fmt.Printf("Number of ways to climb %d stairs: %d\n", steps, result)
}
