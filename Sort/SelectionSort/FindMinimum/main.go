package main

import "fmt"

func main() {
	arrStr := []string{"preya", "sreya", "latha", "pranal", "sravan", "prakashan"}
	var arrInt []int
	arrInt = append(arrInt, 15, 2, 4, 1, 5, 9, 5, 4)
	fmt.Println("the unsorted  string array:", arrStr)
	MinAndSorting(arrStr)
	fmt.Println("the sorted string array:", arrStr)
	fmt.Println("the unsorted  integer array:", arrInt)
	MinForNumber(arrInt)
	fmt.Println("the sorted integer array:", arrInt)

}
func MinAndSorting(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}

		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

}
func MinForNumber(arrint []int) {
	n := len(arrint)
	for i := 0; i < n-1; i++ {
		minIn := i
		for j := i + 1; j < n; j++ {
			if arrint[j] < arrint[minIn] {
				minIn = j
			}
		}
		arrint[i], arrint[minIn] = arrint[minIn], arrint[i]
	}

}