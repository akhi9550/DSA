package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	digits := strings.Split(strconv.Itoa(num), "")
	for i, digit := range digits {
		if digit == "6" {
			digits[i] = "9"
			break
		}
	}
	num, _ = strconv.Atoi(strings.Join(digits, ""))
	return num
}

func main() {
	num := 9669
	fmt.Printf("Original number: %d\n", num)
	updatedNumber := maximum69Number(num)
	fmt.Printf("Updated number: %d\n", updatedNumber)
}
