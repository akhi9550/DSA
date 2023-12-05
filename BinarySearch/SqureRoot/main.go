package main

import "fmt"

func SqureRoot(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	l, r := 1, x
	result := 0
	for l <= r {
		mid := (l + r) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid <= x {
			result = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return result
}
func main() {
	fmt.Println("Square root of 49:", SqureRoot(49))
	fmt.Println("Square root of 8:", SqureRoot(8))
	fmt.Println("Square root of 16:", SqureRoot(16))
	fmt.Println("Square root of 25:", SqureRoot(25))
}
