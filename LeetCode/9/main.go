package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

func main() {
	number := 121
	if isPalindrome(number) {
		fmt.Println(number, "is a palindrome.")
	} else {
		fmt.Println(number, "is not a palindrome.")
	}
}
