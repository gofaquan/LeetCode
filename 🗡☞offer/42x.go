package ___offer


func maxSubArray(nums []int) int {
	m := make([]int, len(nums)+1)
	for i := 0; i < len(nums)-1; i++ {
		m[i+1] = max(m[i]+nums[i+1], nums[i+1])
	}
	return m[len(nums)]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
