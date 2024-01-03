package main

import "fmt"

func main() {
	arr := []int{12, 11, 13, 4, 5, 7}
	fmt.Println("Unsoted Array :-", arr)
	heapSort(arr)
	fmt.Println("Sorted Array:-", arr)
}
func heapSort(arr []int) {
	n := len(arr)
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
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, n, largest)
	}
}
