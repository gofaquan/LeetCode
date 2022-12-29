package ___offer

func numWays(n int) int {

	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}

	return a
}

// 初始 dp
func numWays2(n int) int {
	if n == 0 {
		return 1 //默认？ 我觉得应该是 0....
	}
	res := make([]int, n+1)
	res[0] = 1 // 默认跳 0 级为 1 咯
	res[1] = 1
	for i := 2; i <= n; i++ {
		res[i] = (res[i-1] + res[i-2]) % (1e9 + 7)
	}

	return res[n]
}
