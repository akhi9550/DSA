package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
	n := 5
	for i := 1; i <= n; i++ {
		result := fibonacci(i)
		fmt.Printf("%d ", result)
	}
	pos := fibonacci(n)
	fmt.Println("\nFibonacci of nth position:", pos)

}
