package main

import (
	"fmt"
	"math"
)

func nthUglyNumber(n int) int {
	data := make([]int, n)
	data[0] = 1
	for c2, c3, c5, i := 0, 0, 0, 1; i < n; i++ {
		data[i] = int(math.Min(math.Min(float64(data[c2]*2), float64(data[c3]*3)), float64(data[c5]*5)))
		if data[i] == data[c2]*2 {
			c2++
		}
		if data[i] == data[c3]*3 {
			c3++
		}
		if data[i] == data[c5]*5 {
			c5++
		}
	}
	return data[n-1]
}

func main() {
	n := 10
	fmt.Printf("The %dth ugly number is: %d\n", n, nthUglyNumber(n))
}
