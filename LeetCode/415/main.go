package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	result := ""
	carry := 0
	i, j := len(num1)-1, len(num2)-1

	for i >= 0 || j >= 0 || carry > 0 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			y = int(num2[j] - '0')
			j--
		}

		sum := x + y + carry
		carry = sum / 10
		result = string(sum%10+'0') + result
	}
	return result
}

func main() {
	fmt.Println(addStrings("123", "456"))      // Output: "579"
	fmt.Println(addStrings("11", "123"))       // Output: "134"
	fmt.Println(addStrings("0", "0"))          // Output: "0"
	fmt.Println(addStrings("999", "1"))        // Output: "1000"
	fmt.Println(addStrings("456789", "123456")) // Output: "580245"
}
