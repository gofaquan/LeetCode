package ___offer

func lengthOfLongestSubstring2(s string) (res int) {
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
