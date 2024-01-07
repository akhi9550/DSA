package main

import "fmt"

func main() {
	arr := []int{12, 11, 13, 4, 5, 7, 9}
	fmt.Println("Unsoted Array :-", arr)
	n := len(arr)
	heapSortMax(arr, n)
	fmt.Println("Sorted Array Ascending:-", arr)
	heapSortMin(arr, n)
	fmt.Println("Sorted Array Decending:-", arr)
}

func heapSortMax(arr []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapify(arr, i, 0)
	}
}
func maxHeapify(arr []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	for l < n && arr[l] > arr[largest] {
		largest = l
	}
	for r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, n, largest)
	}
}

func heapSortMin(arr []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		minHeapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		minHeapify(arr, i, 0)
	}
}

func minHeapify(arr []int, n, i int) {
	small := i
	left := 2*i + 1
	right := 2*i + 2
	for left < n && arr[left] < arr[small] {
		small = left
	}
	for right < n && arr[right] < arr[small] {
		small = right
	}
	if small != i {
		arr[i], arr[small] = arr[small], arr[i]
		minHeapify(arr, n, small)
	}
}
