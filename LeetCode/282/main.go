package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	index := 0
	for _, i := range nums {
		if i != 0 {
			nums[index] = i
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println("Original array:", nums)

	moveZeroes(nums)

	fmt.Println("Array after moving zeroes:", nums)
}
