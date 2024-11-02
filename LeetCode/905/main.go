package main

import (
	"fmt"
)

func sortArrayByParity(nums []int) []int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left]%2 == 0 {
			left++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}

	return nums
}

func main() {
	nums := []int{3, 1, 2, 4}
	fmt.Println("Original array:", nums)
	sortedArray := sortArrayByParity(nums)
	fmt.Println("Sorted array by parity:", sortedArray)
}
