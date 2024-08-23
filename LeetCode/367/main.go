package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	i := 0
	for i*i < num {
		i++
	}
	return i*i == num
}

func main() {
	numbers := []int{16, 23, 25, 30, 36, 50}
	for _, num := range numbers {
		if isPerfectSquare(num) {
			fmt.Printf("%d is a perfect square.\n", num)
		} else {
			fmt.Printf("%d is not a perfect square.\n", num)
		}
	}
}
