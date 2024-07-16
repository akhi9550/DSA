package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	result := strStr(haystack, needle)
	fmt.Println("Index:", result)
}
