package leecode

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
