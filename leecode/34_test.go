package leecode

import "sort"

// 自带库的 二分
func _34searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

// 手动实现一个二分
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := SearchFunc(len(nums), func(i int) bool { return nums[i] >= target })
	right := SearchFunc(len(nums), func(i int) bool { return nums[i] >= target+1 }) - 1

	// 前者  保证数组中存在对应的target   	排除类似 [2,2] 3
	// 后者  保证数组中存在对应的target 	排除类似 [2,4] 3
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	return []int{left, right}

}

func SearchFunc(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}
