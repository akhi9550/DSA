package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	arrCopy := append([]int(nil), arr...)
	sort.Ints(arrCopy)
	rank := make(map[int]int)
	for _, n := range arrCopy {
		if _, exists := rank[n]; !exists {
			rank[n] = len(rank) + 1
		}
	}
	for i, n := range arr {
		arr[i] = rank[n]
	}

	return arr
}

func main() {
	arr := []int{40, 10, 20, 30}
	fmt.Println("Original Array:", arr)
	rankedArr := arrayRankTransform(arr)
	fmt.Println("Ranked Array:", rankedArr)
}
