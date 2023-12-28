package main

import "fmt"

func Largest(arr []int, largest, size int) int {
	if size < 0 {
		return largest
	} else {
		if largest < arr[size] {
			largest = arr[size]
		}
	}
	return Largest(arr, largest, size-1)
}
func main() {
	arr := []int{2, 3, 4, 53, 8, 56, 9}
	largest := arr[0]
	size := len(arr) - 1
	result := Largest(arr, largest, size)
	fmt.Println("Largest Element is:", result)
}
