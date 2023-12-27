package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(i int) bool {
	str := strconv.Itoa(i)
	if len(str) <= 1 {
		return true
	}
	if str[0] != str[len(str)-1] {
		return false
	}

	v, _ := strconv.Atoi(str[1 : len(str)-1])
	return isPalindrome(v)
}
func main() {
	String := 121
	String1 := 21
	result := isPalindrome(String)
	if result {
		fmt.Printf("%d is palindrome", String)
	} else {
		fmt.Printf("%d is not palindrome", String)
	}
	result1 := isPalindrome(String1)
	if result1 {
		fmt.Printf("\n %d is palindrome", String1)
	} else {
		fmt.Printf("\n %d is not palindrome", String1)
	}
}
