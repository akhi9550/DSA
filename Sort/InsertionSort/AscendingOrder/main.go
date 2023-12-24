package main

import "fmt"
func InsertionSort(arr []int, n int) []int {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] >= key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	arr := []int{2, 43, 32, 54, 3, 1, 98}
	n := len(arr)
	fmt.Println("Unsorted Array is :", arr)
	result := InsertionSort(arr, n)
	fmt.Println("Sorted Array is :", result)
}
