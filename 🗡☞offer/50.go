package ___offer

func firstUniqChar(s string) byte {
	m := map[uint8]int{}
	for j := 0; j < len(s); j++ {
		m[s[j]]++
	}

	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return s[i]
		}
	}

	return ' '
}
