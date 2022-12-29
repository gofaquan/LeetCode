package ___offer

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	if n == 0 {
		return 0
	}

	a, b := 1, 1
	for i := 1; i < n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}

	return a
}

// 初始 dp
func fib2(n int) int {
	if n == 0 {
		return 0
	}
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	for i := 2; i <= n; i++ {
		res[i] = (res[i-1] + res[i-2]) % (1e9 + 7)
	}

	return res[n]
}
