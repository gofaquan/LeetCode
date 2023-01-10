package ___offer

func f(n int, m int) int {
	if n == 0 {
		return 0
	}
	x := f(n-1, m)
	return (m + x) % n
}

func lastRemaining(n int, m int) int {
	return f(n, m)
}
