package ___offer

import "strconv"

func minNumber(nums []int) string {
	strSlice := make([]string, len(nums))
	for i, num := range nums {
		strSlice[i] = strconv.Itoa(num)
	}

	for i := 0; i < len(strSlice)-1; i++ {
		if (strSlice[i+1] + strSlice[i]) < (strSlice[i] + strSlice[i+1]) {
			strSlice[i+1], strSlice[i] = strSlice[i], strSlice[i+1]
		}
	}

	var res string
	for _, s := range strSlice {
		res += s
	}
	return res
}
