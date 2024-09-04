package main

import (
	"fmt"
)

func isBadVersion(version int) bool {
	return version >= 4
}

func firstBadVersion(n int) int {
	i, j := 1, n
	for i + 1 < j {
		mid := (i + j) / 2
		if isBadVersion(mid) {
			j = mid
		} else {
			i = mid
		}
	}
	if isBadVersion(i) {
		return i
	}
	return j
}

func main() {
	n := 10
	firstBad := firstBadVersion(n)
	fmt.Printf("The first bad version is: %d\n", firstBad)
}
