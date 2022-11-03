package leecode

// 朴素做法，嗯写
func strStr1(haystack string, needle string) int {
	//  遍历到 needle的最后一个字符正好是 haystack 的最后一个字符，所以是 相减
	for i := 0; i < len(haystack)-len(needle); i++ {
		j := i
		k := 0
		for ; k < len(needle) && haystack[j] == needle[k]; j, k = j+1, k+1 {
		}
		if k == len(needle) {
			return i
		}
	}

	return -1
}

func strStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}

	pi := make([]int, m)
	// 拿到 pi 数组
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}

		if needle[i] == needle[j] {
			j++
		}

		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}

		if j == m {
			return i - m + 1
		}
	}

	return -1
}
