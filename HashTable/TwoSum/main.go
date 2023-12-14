package main

import "fmt"

func main() {
	arr := []int{2, 4, 5, 9, 8, 6}
	target := 17
	result := twoSum(arr, target)
	fmt.Printf("The target %d is the sum of %d", target, result)
}
func twoSum(arr []int, target int) []int {
	table := make(map[int]int)
	for i, v := range arr {
		result := target - v
		if j, found := table[result]; found {
			return []int{j, i}
		}
		table[v] = i
	}
	return []int{}
}
