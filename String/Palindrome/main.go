package main

import "fmt"

func Palindrome(str string) {
	c := 0
	for i, j := 0, len(str)-1; i <= len(str)/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			c++
		}
	}
	if c == 0 {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}
}
func main() {
	Palindrome("MALAYALAM")
	Plaindkjd("malayALAm")
}
func Plaindkjd(str string) {
	c := 0
	char := []rune(str)
	for i, j := 0, len(str)-1; i <= len(str)/2; i, j = i+1, j-1 {
		if char[i] != char[j] && char[i] != char[j]-32 && char[i] != char[j]+32 {
			c++
		}
	}
	if c == 0 {
		fmt.Println("palindrome")
	}
}

// func main() {
// 	Palindrome("malayalam")
// 	// PalindromeWithoutCapitalOsSmall("12321")
// }
// func Palindrome(str string) {
// 	count := 0
// 	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
// 		if str[i] != str[j] {
// 			count = 1
// 		}
// 	}
// 	if count == 1 {
// 		fmt.Println("Not a Palindrome")
// 	} else {
// 		fmt.Println("Palindrome")
// 	}
// }

// func PalindromeWithoutCapitalOsSmall(str string) {
// 	character := []rune(str)
// 	count := 0
// 	for i, j := 0, len(character)-1; i < len(character)/2; i, j = i+1, j-1 {
// 		if character[i] != character[j] && character[i] != character[j]-32 && character[i] != character[j]+32 {
// 			count = 1
// 		}
// 	}
// 	if count == 0 {
// 		fmt.Println("Palindrome")
// 	} else {
// 		fmt.Println("Not a Palindrome")
// 	}
// }
