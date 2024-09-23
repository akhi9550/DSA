package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func findDisappearedNumbers(nums []int) []int {
	for _, val := range nums {
		i := abs(val) - 1
		nums[i] = -1 * abs(nums[i])
	}

	result := make([]int, 0)
	for i, val := range nums {
		if val > 0 {
			result = append(result, i+1)
		}
	}

	return result
}
func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := findDisappearedNumbers(nums)
	fmt.Println("Disappeared numbers:", result)
}
