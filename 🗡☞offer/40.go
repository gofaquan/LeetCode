package ___offer

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	var res []int
	res = append(res, arr[:k]...)
	return res
}
