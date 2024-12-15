package main

import (
	"fmt"
	"sort"
)

func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	return nums
}

func main() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	fmt.Println("Original Slice:", nums)
	result := numberGame(nums)
	fmt.Println("Modified Slice:", result)
}
