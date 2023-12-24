package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum := Recursion(arr)
	fmt.Println("Sum is :", sum)
}
func Recursion(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + Recursion(arr[1:])
}

// // otherMethod
// func Recursionn(arr []int, index int) int {
// 	if index == len(arr) {
// 		return 0
// 	}
// 	return arr[index] + Recursionn(arr, index+1)
// }
