package main

import (
	"fmt"
)

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return isPalindrome(s[l+1:r+1]) || isPalindrome(s[l:r])
		}
		l, r = l+1, r-1
	}

	return true
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l, r = l+1, r-1
	}

	return true
}

func main() {
	tests := []string{
		"abca",
		"racecar",
		"deified",
		"hello",
	}

	for _, test := range tests {
		if validPalindrome(test) {
			fmt.Printf("%s can be a palindrome after removing at most one character.\n", test)
		} else {
			fmt.Printf("%s cannot be a palindrome even after removing one character.\n", test)
		}
	}
}
