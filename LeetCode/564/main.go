package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n string
	fmt.Println("Enter a number:")
	fmt.Scan(&n)
	
	result := nearestPalindromic(n)
	fmt.Println("Nearest Palindromic number:", result)
}

func nearestPalindromic(n string) string {
	l := len(n)
	num, _ := strconv.Atoi(n)
	if num == 0 {
		return "-1"
	}

	a, _ := strconv.Atoi(n[:(l+1)/2])
	l2 := len(strconv.Itoa(a))
	b := a + 1
	c := a - 1

	var res int
	diff := math.MaxInt64
	for _, v := range []int{b, a, c} {
		var pali int
		if len(strconv.Itoa(v)) > l2 {
			if l%2 == 1 {
				v /= 10
			}
			pali = makePali(v, 1-l%2)
		} else if len(strconv.Itoa(v)) < l2 || v == 0 {
			if l%2 != 1 {
				v = v*10 + 9
			}
			pali = makePali(v, 1-l%2)
		} else {
			pali = makePali(v, l%2)
		}
		if pali == num {
			continue
		}
		if abs(pali-num) <= diff {
			res = pali
			diff = abs(pali - num)
		}
	}
	return strconv.Itoa(res)
}

func makePali(a int, flag int) int {
	b := a
	if flag == 1 {
		b /= 10
	}
	for b != 0 {
		a = a*10 + b%10
		b /= 10
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
