package main

import (
	"fmt"
)

func xorQueries(arr []int, queries [][]int) []int {
	result := make([]int, len(queries))
	if len(arr) == 0 || len(queries) == 0 {
		return result
	}

	xor := make([]int, len(arr))
	xor[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		xor[i] = arr[i] ^ xor[i-1]
	}

	for idx, query := range queries {
		if query[0] == 0 {
			result[idx] = xor[query[1]]
		} else {
			result[idx] = xor[query[1]] ^ xor[query[0]-1]
		}
	}

	return result
}

func main() {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{
		{0, 1},
		{1, 2},
		{0, 3},
		{3, 3},
	}

	result := xorQueries(arr, queries)

	fmt.Println("Results:", result)
}
