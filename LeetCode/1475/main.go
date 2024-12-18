package main

import (
	"fmt"
)

func finalPrices(prices []int) []int {
	stack := []int{}
	for i, v := range prices {
		for len(stack) != 0 && prices[stack[len(stack)-1]] >= v {
			prices[stack[len(stack)-1]] -= v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return prices
}

func main() {
	prices := []int{8, 4, 6, 2, 3}
	finalPricesWithDiscount := finalPrices(prices)
	fmt.Println("Original prices: [8, 4, 6, 2, 3]")
	fmt.Println("Final prices after discount:", finalPricesWithDiscount)
}
