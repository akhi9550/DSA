package main

import "fmt"

func pivotIndex(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	sum := 0
	for _, x := range nums {
		sum += x
	}

	leftSum := 0
	for i, x := range nums {
		rightSum := sum - x - leftSum
		if leftSum == rightSum {
			return i
		}
		leftSum += x
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println("Pivot index:", pivotIndex(nums)) 

	nums = []int{1, 2, 3}
	fmt.Println("Pivot index:", pivotIndex(nums)) 

	nums = []int{2, 1, -1}
	fmt.Println("Pivot index:", pivotIndex(nums)) 
}
