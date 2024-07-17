package main

import "fmt"

func search(nums []int, target int) int {
	findPivot := func(nums []int) int {
		left := 0
		right := len(nums) - 1

		for left < right {
			mid := left + (right-left)/2

			if nums[mid] > nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		}

		return left
	}

	pivot := findPivot(nums)
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		midPivot := (mid + pivot) % len(nums)

		if nums[midPivot] == target {
			return midPivot
		} else if nums[midPivot] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	result := search(nums, target)
	if result != -1 {
		fmt.Printf("Target %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Target %d not found in the array\n", target)
	}
}
