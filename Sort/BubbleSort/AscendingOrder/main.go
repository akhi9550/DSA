package main

import "fmt"

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		flag := 0
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
	return arr
}
func main() {
	arr := []int{32, 56, 343, 64, 2}
	fmt.Println("Unsorted Arry is :", arr)
	result := BubbleSort(arr)
	fmt.Println("Sorted Array is :", result)
}
