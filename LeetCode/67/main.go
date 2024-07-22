package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func addBinary(a string, b string) string {
	carry := 0
	A := []rune(a)
	B := []rune(b)
	slices.Reverse(A)
	slices.Reverse(B)
	res := []rune{}

	for i := 0; i < len(A) || i < len(B); i++ {
		sum := 0
		if i < len(A) && A[i] == '1' {
			sum++
		}
		if i < len(B) && B[i] == '1' {
			sum++
		}
		sum += carry
		if sum == 3 {
			res = append(res, '1')
			carry = 1
		} else if sum == 2 {
			res = append(res, '0')
			carry = 1
		} else if sum == 1 {
			res = append(res, '1')
			carry = 0
		} else {
			res = append(res, '0')
			carry = 0
		}
	}
	if carry > 0 {
		res = append(res, '1')
	}
	slices.Reverse(res)
	r := string(res)
	return r
}

func main() {
	a := "1010"
	b := "1011"
	result := addBinary(a, b)

	fmt.Println("Binary addition of", a, "and", b, "is:", result)
}
