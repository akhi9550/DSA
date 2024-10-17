package main

import (
	"fmt"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	h := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	result := ""
	for num != 0 {
		result = string(h[num&15]) + result
		num = int(uint32(num) >> 4)
	}
	return result
}

func main() {
	num := 26
	hex := toHex(num)
	fmt.Printf("The hexadecimal representation of %d is: %s\n", num, hex)

	num = -1
	hex = toHex(num)
	fmt.Printf("The hexadecimal representation of %d is: %s\n", num, hex)
}
