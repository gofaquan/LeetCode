package ___offer

import "math"

// 求出将长度为 n 的绳子剪成若干段后，各段长度的最大乘积
func cuttingRope(n int) int {
	// dp[i] 表示将长度为 i 的绳子剪成若干段后，各段长度的最大乘积
	dp := make([]int, n+1)
	// 长度为 1 和 2 的绳子无法剪成若干段
	dp[1] = 1
	dp[2] = 1

	// 从长度为 3 开始，逐个求出长度为 i 的绳子的最优解
	for i := 3; i <= n; i++ {
		// curMax 表示将长度为 i 的绳子剪成若干段后，各段长度的最大乘积
		curMax := 0
		// 枚举所有可能的割点，求出长度为 i 的绳子的最优解
		for j := 1; j < i; j++ {
			// 比较把绳子剪成 j 和 i-j 两段的乘积，和把绳子剪成 j 和 dp[i-j] 两段的乘积，取最大值
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}

// 求出两个整数的最大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}

	a, b := n/3, n%3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	}
	if b == 1 {
		return int(math.Pow(3, float64(a-1)) * 4)
	}
	return int(math.Pow(3, float64(a)) * 2)
}
