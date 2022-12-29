package ___offer

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}

}

// 初始 dp
func maxSubArray(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(f[i-1]+nums[i], nums[i])
	}
	ans := f[0]
	for i := 1; i < len(nums); i++ {
		ans = max(ans, f[i])
	}
	return ans
}

// 滚动数组 dp
func maxSubArray(nums []int) int {
	n := len(nums)
	f := make([]int, 2)
	f[0] = nums[0]
	ans := f[0]
	for i := 1; i < n; i++ {
		f[i%2] = max(f[(i-1)%2]+nums[i], nums[i])
		ans = max(ans, f[i%2])
	}
	return ans
}

// 滚动数组优化 dp
func maxSubArray(nums []int) int {
	n := len(nums)
	m := nums[0] // 一个 m 就够了
	ans := m
	for i := 1; i < n; i++ {
		m = max(m+nums[i], nums[i])
		ans = max(ans, m)
	}
	return ans
}

// 滚动数组再优化 dp ， 也就是 leetcode 官方题解
func maxSubArray(nums []int) int {
	n := len(nums)
	ans := nums[0]
	for i := 1; i < n; i++ {
		// 不用变量，直接在数组上操作
		nums[i] = max(nums[i-1]+nums[i], nums[i])
		ans = max(ans, nums[i])
	}
	return ans
}
