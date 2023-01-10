package ___offer

func findContinuousSequence(target int) [][]int {
	var res [][]int
	for i, j := 1, 2; i < j; {
		cur := (j + i) * (j - i + 1) / 2
		if cur == target {
			var ans []int
			for x := i; x <= j; x++ {
				ans = append(ans, x)
			}
			res = append(res, ans)
			// 记得后推
			i++
		} else if cur < target {
			j++
		} else {
			i++
		}

	}

	return res
}
