package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	res := 0
	i, j := 0, 0

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			res++
			i++
			j++
		} else {
			j++
		}
	}
	return res
}

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))    // Output: 1
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))    // Output: 2
	fmt.Println(findContentChildren([]int{10, 9, 8, 7}, []int{5, 6})) // Output: 0
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{3}))        // Output: 1
	fmt.Println(findContentChildren([]int{5, 10, 2}, []int{4, 1, 7})) // Output: 2
}
