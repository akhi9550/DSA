package main

import (
	"fmt"
	"sort"
	"strconv"
)

func lexicalOrder(n int) []int {
	numbers := make([]string, n)
	for i := 0; i < n; i++ {
		numbers[i] = strconv.Itoa(i + 1)
	}
	sort.Strings(numbers)

	result := make([]int, n)
	for i, number := range numbers {
		result[i], _ = strconv.Atoi(number)
	}

	return result
}

func main() {
	n := 13
	lexicalNumbers := lexicalOrder(n)
	fmt.Println("Lexicographical Order of numbers 1 to", n, ":")
	fmt.Println(lexicalNumbers)
}
