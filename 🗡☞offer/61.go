package ___offer

import "sort"

func isStraight(nums []int) bool {
	// 统计数组中 0 的个数
	zeroCount := 0
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		}
	}

	// 对数组进行排序
	sort.Ints(nums)

	// 遍历数组，检查是否有重复的数字
	for i := zeroCount; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return false
		}
	}

	// 遍历数组，计算最大值和最小值的差值
	maxVal := nums[len(nums)-1]
	minVal := nums[zeroCount]
	diff := maxVal - minVal

	// 如果最大值和最小值的差值小于等于 4，就返回 true，否则返回 false
	if diff <= 4 {
		return true
	}
	return false
}
