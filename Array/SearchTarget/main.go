package main

import "fmt"

func Sum(arr []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				result = (append(result, []int{arr[i], arr[j]}))
				return result
			}
		}
	}
	return result
}
func main() {
	arr := []int{10, 23, 43, 43, 56, 9}
	target := 33
	result := Sum(arr, target)
	for i := 0; i < len(result); i++ {
		fmt.Println("result is", result[i])
	}
}
