package main

import (
	"fmt"
)

func BinarySearch(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] < target {
		return BinarySearch(arr, target, mid+1, right)
	} else {
		return BinarySearch(arr, target, left, mid-1)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	target := 8
	result := BinarySearch(arr, target, 0, len(arr)-1)
	if result != -1 {
		fmt.Printf("Element %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
