package main

import (
	"fmt"
)

func LinearSearch(arr []int, target int) int {
	for i := range arr {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
func main() {
	arr := []int{1, 3, 5, 43, 56, 76, 45}
	target := 1
	result := LinearSearch(arr, target)
	if result != -1 {
		fmt.Printf("Element %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Element %d not found\n", target)

	}
}
