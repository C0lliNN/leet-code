package main

import (
	"fmt"
	"math"
)

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit1(prices1))
	fmt.Println(maxProfit1(prices2))

	fmt.Println(maxProfit2(prices1))
	fmt.Println(maxProfit2(prices2))
}

func maxProfit1(prices []int) int {
	maximumProfit := 0

	for i := 0; i < len(prices)-1; i++ {
		for j := i; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maximumProfit {
				maximumProfit = profit
			}
		}
	}

	return maximumProfit
}

func maxProfit2(prices []int) int {
	maxProfit, minPrice := 0, math.MaxInt
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		profit := price - minPrice

		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
