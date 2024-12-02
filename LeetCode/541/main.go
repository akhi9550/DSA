package main

import (
	"fmt"
)

func reverseStr(s string, k int) string {
	s1 := []rune(s)
	i := 0
	for i < len(s1) {
		end := i + k
		if end > len(s1) {
			end = len(s1)
		}
		for i2, j := i, end-1; i2 < j; i2, j = i2+1, j-1 {
			s1[i2], s1[j] = s1[j], s1[i2]
		}
		i += 2 * k
	}
	return string(s1)
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))    // Output: "bacdfeg"
	fmt.Println(reverseStr("abcd", 4))       // Output: "dcba"
	fmt.Println(reverseStr("abcdefghij", 3)) // Output: "cbadefihgj"
	fmt.Println(reverseStr("a", 2))          // Output: "a"
	fmt.Println(reverseStr("abcdef", 6))     // Output: "fedcba"
}
