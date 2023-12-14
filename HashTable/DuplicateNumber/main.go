package main

import "fmt"

func main() {
	arr := []int{2, 3, 2, 3, 4, 6, 5, 7, 5}
	fmt.Println("Array before deleting the duplicate", arr)
	result := DuplicateDelete(arr)
	fmt.Println("Array after deleting the duplicate", result)
}
func DuplicateDelete(arr []int) []int {
	table := make(map[int]bool)
	var Array []int
	for _, v := range arr {
		if !table[v] {
			Array = append(Array, v)
			table[v] = true
		}
	}
	return Array
}
