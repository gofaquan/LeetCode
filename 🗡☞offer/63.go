package ___offer

import "math"

func maxProfit(prices []int) int {
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
