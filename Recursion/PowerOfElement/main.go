package main

import "fmt"

func Power(n, x int) int {
	if x == 0 {
		return 1
	}
	return n * Power(n, x-1)
}
func main() {
	n, x := 2, 5
	result := Power(n, x)
	fmt.Println("Power is:", result)
}
