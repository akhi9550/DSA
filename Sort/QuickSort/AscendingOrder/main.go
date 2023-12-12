package main

import "fmt"

func QuickSort(arr []int, start, end int) []int {
	if start < end {
		loc := Partition(arr, start, end)
		QuickSort(arr, start, loc-1)
		QuickSort(arr, loc+1, end)
	}
	return arr
}

func Partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start - 1
	for j := start; j <= end-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{1, 5, 3, 8, 99, 56, 45, 34, 7, 9}
	fmt.Println("Unsorted Array is:", arr)
	start := 0
	n := len(arr) - 1
	end := n
	result := QuickSort(arr, start, end)
	fmt.Println("Sorted Array is :", result)
}
