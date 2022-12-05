package leecode

func _3lengthOfLongestSubstring(s string) int {
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

func _3lengthOfLongestSubstring2(s string) (res int) {
	m := map[uint8]int{}
	for i, j := 0, 0; i < len(s); i++ {
		m[s[i]]++
		for m[s[i]] > 1 {
			m[s[j]]--
			j++
		}

		if res < i-j+1 {
			res = i - j + 1
		}
	}

	return
}
