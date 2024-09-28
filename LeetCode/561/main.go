package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	var result int
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		result += nums[i] 
	}
	return result
}

func main() {
	nums := []int{1, 4, 3, 2}
	fmt.Println("Array Pair Sum:", arrayPairSum(nums))
}
