package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var profit = 0
	var minPrice = prices[0]
	
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > profit {
			profit = prices[i] - minPrice
		}
	}
	
	return profit
}

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Example 1 - Input: %v, Output: %d\n", prices1, maxProfit(prices1))

	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Printf("Example 2 - Input: %v, Output: %d\n", prices2, maxProfit(prices2))
}
