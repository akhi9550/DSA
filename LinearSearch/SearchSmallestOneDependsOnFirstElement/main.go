package main

import "fmt"

func SmallestElement(arr []int) int {
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[0] > arr[i] {
			arr[0] = arr[i]
			index = i
		}
	}
	return index
}

func main() {
	arr := []int{6, 78, 65, 4, 2, 98}
	result := SmallestElement(arr)
	if result == -1 {
		fmt.Println("Element Not Found Smallest is that one")
	} else {
		fmt.Println("Element Found at index", result)
	}
}
