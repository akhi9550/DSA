package main

import "fmt"

func main() {
	arr := []int{2, 5, 8, 45, 45, 45, 2, 2, 6, 6, 7}
	fmt.Println("Before Delete Array is :", arr)
	Result := removeDuplicates(arr)
	fmt.Println("Count of Duplicates :", Result)
}
func removeDuplicates(arr []int) int {
	table := make(map[int]int)
	count:=0
	for _, v := range arr {
		table[v]++
		if table[v]==2 {
			count++
		}
	}
	return count
}
