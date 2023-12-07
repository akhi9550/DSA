package main

import "fmt"

func LinearSearch(arr []int, target int) []int {
	var ar []int
	for i := range arr {
		if arr[i] == target {
			ar = append(ar, i)
		}
	}
	return ar
}


func main() {
	arr := []int{1, 2, 32, 3, 32, 43, 32, 8}
	target := 32
	result := LinearSearch(arr, target)
	if len(result) > 0 {
		fmt.Printf("Element %d Found at these position %d", target, result)
	} else {
		fmt.Printf("Element %d Not Found", target)
	}

}
