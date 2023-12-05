package main

import (
	"fmt"
	"math"
)

func Closest(arr []int, target int) int {
	low, high := 0, len(arr)-1
	close := math.MaxInt16
	closestIndex := -1
	for low <= high {
		mid := (low + high) / 2
		if abs(arr[mid]-target) < abs(close-target) {
			closestIndex = mid
			close = arr[mid]
		}
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		}
		if arr[mid] > target {
			high = mid - 1
		}
	}
	return closestIndex
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	arr := []int{2, 3, 4, 54, 55, 65, 67, 86, 98}
	target := 70
	closestIndex := Closest(arr, target)
	if closestIndex != -1 {
		fmt.Printf("Closest element to %d is %d at index %d\n", target, arr[closestIndex], closestIndex)
	} else {
		fmt.Println("Array is empty")
	}
}
