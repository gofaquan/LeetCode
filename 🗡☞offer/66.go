package ___offer

// constructArr 返回整型数组 a 的一个副本，其中第 i 个元素表示除了下标 i 以外的元素的积。
func constructArr(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	// 创建一个空数组用于保存结果。
	b := make([]int, len(a))
	// 将结果数组的第一个元素设为 1。
	b[0] = 1
	// 遍历数组 a，计算当前元素之前的所有元素的积。
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}
	// 初始化一个变量 p，用于保存当前元素之后的所有元素的积。
	p := 1
	// 倒序遍历数组 a，计算当前元素之后的所有元素的积。
	for i := len(a) - 1; i >= 0; i-- {
		// 将当前元素的结果乘上之后的所有元素的积。
		b[i] *= p
		// 更新之后的所有元素的积。
		p *= a[i]
	}
	// 返回结果数组。
	return b
}
