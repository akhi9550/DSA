package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func main() {
	arr := []int{1, 21, 34, 43, 56, 72, 87, 90}
	target := 72
	result := BinarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element %d Found at Position %d", target, result)
	} else {
		fmt.Printf("Element %d Not Found", target)
	}
}
