package main

import "fmt"

func searchElement(nums []int, target int) int {
	low, high := 0, len(nums)
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	target1 := 11
	target3 := 1
	target4 := 0

	result1 := searchElement(arr, target1)
	result3 := searchElement(arr, target3)
	result4 := searchElement(arr, target4)

	fmt.Printf("Target %d should be inserted at index %d\n", target1, result1)
	fmt.Printf("Target %d should be inserted at index %d\n", target3, result3)
	fmt.Printf("Target %d should be inserted at index %d\n", target4, result4)
}
