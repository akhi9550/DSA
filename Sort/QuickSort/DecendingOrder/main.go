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

	for i <= j {
		for i <= h && arr[i] >= pivot {
			i++
		}
		for arr[j] < pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func main() {
	arr := []int{1, 5, 3, 8, 99, 56, 45, 34, 7, 9}
	fmt.Println("Unsorted Array is:", arr)
	start := 0
	n := len(arr) - 1
	end := n
	result := QuickSort(arr, start, end)
	fmt.Println("Sorted Array in Descending Order is:", result)
}
