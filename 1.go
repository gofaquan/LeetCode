package main

import "math"

type Node struct {
	Val      int
	Children []*Node
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
