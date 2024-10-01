package main

import (
	"fmt"
)

func canArrange(arr []int, k int) bool {
	count := make([]int, k)

	for _, n := range arr {
		grp := n % k
		grp = (grp + k) % k
		count[grp]++
	}

	if count[0]%2 != 0 {
		return false
	}

	for i := 1; i < k; i++ {
		j := k - i
		if count[i] != count[j] {
			return false
		}
	}

	return true
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	k := 5

	if canArrange(arr, k) {
		fmt.Println("The array can be arranged.")
	} else {
		fmt.Println("The array cannot be arranged.")
	}
}
