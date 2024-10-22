package main

import (
	"fmt"
)

func removeDigit(number string, digit byte) string {
	var r string
	for i := 0; i < len(number); i++ {
		if number[i] != digit {
			continue
		}

		n := number[:i] + number[i+1:]
		if n > r {
			r = n
		}
	}
	return r
}

func main() {
	number := "123523"
	digit := byte('2')
	result := removeDigit(number, digit)
	fmt.Println("Original number:", number)
	fmt.Println("Digit to remove:", string(digit))
	fmt.Println("Largest possible number after removing digit:", result)
}
