package ___offer

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt
	profit := make([]int, len(prices)+1)
	for i, price := range prices {
		if price < min {
			min = price
		}

		profit[i+1] = max(profit[i], price-min)
	}

	return profit[len(prices)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit2(prices []int) int {
	max, min := 0, math.MaxInt

	for _, price := range prices {
		if price < min {
			min = price
		}
		if price-min > max {
			max = price - min
		}
	}

	return max
}
