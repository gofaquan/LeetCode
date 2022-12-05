package leecode

//  位操作统计二进制中 1 的个数

func _15hammingWeight(num uint32) (ans int) {
	for ; num > 0; num = num >> 1 {
		if num&1 == 1 {
			ans++
		}
	}

	return
}

func _15hammingWeight2(num uint32) int {
	res := 0
	for num > 0 {
		num -= lowbit(num)
		res++
	}

	return res
}

// 返回最后一位数字
func lowbit(x uint32) uint32 {
	return x & -x
}

func _15hammingWeight3(num uint32) int {
	res := 0
	for num > 0 {
		num = num & (num - 1)
		res++
	}

	return res
}
