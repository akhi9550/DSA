package main

import (
	"fmt"
)

func missingInteger(nums []int) int {
	table := make(map[int]bool)
	for _, v := range nums {
		table[v] = true
	}
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 != nums[i] {
			break
		}
		sum += nums[i]
	}
	for table[sum] {
		sum++
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 5}
	fmt.Println("Missing integer:", missingInteger(nums))
}
