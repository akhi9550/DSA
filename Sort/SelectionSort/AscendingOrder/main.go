package main

import "fmt"

func SelectionSort(arr []int, n int) []int {
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

func main() {
	arr := []int{2, 43, 32, 54, 3, 1, 98}
	n := len(arr)
	fmt.Println("Unsorted Array is :", arr)
	result := SelectionSort(arr, n)
	fmt.Println("Sorted Array is :", result)
}
