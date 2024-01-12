package main

import "fmt"

func MergeSort(arr []int, start, end int) {
	if start < end {
		mid := (start + end) / 2
		MergeSort(arr, start, mid)
		MergeSort(arr, mid+1, end)
		Merge(arr, start, mid, end)
	}
}

func Merge(arr []int, start, mid, end int) {
	i := start
	j := mid + 1
	var b []int
	for i <= mid && j <= end {
		if arr[i] >= arr[j] { 
			b = append(b, arr[i])
			i++
		} else {
			b = append(b, arr[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		b = append(b, arr[i])
	}
	for ; j <= end; j++ {
		b = append(b, arr[j])
	}
	for i := start; i <= end; i++ {
		arr[i] = b[i-start]
	}
}

func main() {
	arr := []int{2, 43, 54, 6, 34, 77, 45, 23, 87, 99, 7, 5}
	start := 0
	end := len(arr) - 1
	MergeSort(arr, start, end)
	fmt.Println("Sorted Array in Descending Order is:", arr)
}
