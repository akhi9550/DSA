package main

import (
	"fmt"
)

func lemonadeChange(bills []int) bool {
	var m = make(map[int]int) 
	for _, val := range bills {
		m[val]++
		if val == 10 {
			if m[5] < 1 {
				return false
			} else {
				m[5]--
			}
		}
		if val == 20 {
			if m[10] > 0 && m[5] > 0 {
				m[10]--
				m[5]--
			} else if m[5] >= 3 {
				m[5] -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	bills1 := []int{5, 5, 5, 10, 20}
	bills2 := []int{5, 5, 10, 10, 20}
	bills3 := []int{5, 10, 5, 20, 5, 10, 20, 5, 20}

	fmt.Println("Test Case 1:", lemonadeChange(bills1)) // Expected: true
	fmt.Println("Test Case 2:", lemonadeChange(bills2)) // Expected: false
	fmt.Println("Test Case 3:", lemonadeChange(bills3)) // Expected: false
}
