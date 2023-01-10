package ___offer

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// majorityElement 返回整型数组中出现次数大于数组长度一半的数。
// 如果不存在这样的数，函数返回 -1。
func majorityElement2(nums []int) int {
	// candidate 是当前的候选数字。
	candidate := -1
	// count 是当前候选数字的出现次数。
	count := 0
	// 遍历整型数组。
	for _, num := range nums {

		if count == 0 {
			candidate = num
		}

		// 如果当前数字等于候选数字，则计数器加一。
		if num == candidate {
			count++
			// 如果计数器为零，则更新候选数字并将计数器设为一。
		} else {
			count--
		}
	}
	// 返回候选数字。
	return candidate
}
