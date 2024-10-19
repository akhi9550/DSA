package main

import "fmt"

func findK(k int) int {
	if k <= 1 {
		return 0
	}
	square := 2
	for square < k {
		square *= 2
	}
	return 1 - findK(square - k)
}

func findKthBit(k int) byte { 
	if findK(k) == 0 {
		return '0'
	}
	return '1'
}

func main() {
	var k int
	fmt.Println("Enter the value of k:")
	fmt.Scanf("%d", &k)
	result := findKthBit(k)
	fmt.Printf("The %d-th bit is: %c\n", k, result)
}
