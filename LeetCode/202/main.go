package main

import (
	"fmt"
)

func isHappy(n int) bool {
	visited := make(map[int]bool)
	for n != 1 {
		if _, ok := visited[n]; ok {
			return false
		}
		visited[n] = true
		newnum := 0
		for n > 0 {
			v := n % 10
			newnum += v * v
			n = n / 10
		}
		n = newnum
	}
	return true
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isHappy(num) {
		fmt.Printf("%d is a happy number\n", num)
	} else {
		fmt.Printf("%d is not a happy number\n", num)
	}
}
