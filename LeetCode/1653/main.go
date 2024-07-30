package main

import (
	"fmt"
)

func minimumDeletions(s string) int {
	aCounts := 0
	for _, c := range s {
		if c == 'a' {
			aCounts++
		}
	}

	bCountsLeft := 0
	res := len(s)
	for _, c := range s {
		if c == 'a' {
			aCounts--
		}
		res = min(res, bCountsLeft+aCounts)
		if c == 'b' {
			bCountsLeft++
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var s string
	fmt.Println("Enter the string:")
	fmt.Scan(&s)

	result := minimumDeletions(s)
	fmt.Printf("Minimum deletions needed to make the string balanced: %d\n", result)
}
