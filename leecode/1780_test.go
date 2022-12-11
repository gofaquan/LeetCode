package leecode

func checkPowersOfThree(n int) bool {
	for n != 0 {
		if n%3 == 0 || n%3 == 1 {
			n /= 3
		} else {
			return false
		}
	}

	return true
}
