package leecode

func lengthOfLongestSubstring(s string) int {
	var a = make([]int, len(s))
	res := 0

	for i, j, l := 0, 0, len(s); i < l; i++ {
		a[s[i]]++
		for a[s[i]] > 1 {
			a[s[j]]--
			j++
		}

		if res < i-j+1 {
			res = i - j + 1
		}
	}
	return res
}
