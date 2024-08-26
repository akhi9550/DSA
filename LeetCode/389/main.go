package main

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	var m byte
	i := 0
	for i = 0; i < len(s); i++ {
		m += t[i]
		m -= s[i]
	}
	return m + t[i]
}

func main() {
	s := "abcd"
	t := "abcde"
	result := findTheDifference(s, t)
	fmt.Printf("The difference is: %c\n", result)
}
