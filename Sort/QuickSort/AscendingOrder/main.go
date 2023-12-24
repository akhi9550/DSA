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
func Partition(arr []int, l, h int) int {
	pivot := arr[l]
	i := l + 1
	j := h
	for i < j {
		for arr[i] <= pivot && i < h {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}
// func partition(arr []int, l, h int) int {
// 	pivot := arr[h] // Change pivot to the last element
// 	i := l - 1
// 	for j := l; j < h; j++ {
// 		if arr[j] <= pivot {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 		}
// 	}
// 	arr[i+1], arr[h] = arr[h], arr[i+1]
// 	return i + 1
// }

func main() {
	arr := []int{1, 5, 3, 8, 99, 56, 45, 34, 7, 9}
	fmt.Println("Unsorted Array is:", arr)
	start := 0
	n := len(arr) - 1
	end := n
	result := QuickSort(arr, start, end)
	fmt.Println("Sorted Array is :", result)
}
