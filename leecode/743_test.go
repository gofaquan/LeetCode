package leecode

import "math"

func networkDelayTime(times [][]int, n int, k int) (ans int) {
	const inf = math.MaxInt64 / 2
	dis := make([]int, n)
	for i := 0; i < len(dis); i++ {
		dis[i] = inf
	}
	dis[k-1] = 0
	g := make([][]int, n)

	for i := 0; i < len(g); i++ {
		g[i] = make([]int, n)
		for j := 0; j < len(g[i]); j++ {
			g[i][j] = inf
		}
	}

	for _, time := range times {
		g[time[0]-1][time[1]-1] = time[2]
	}

	used := make([]bool, n)
	for i := 0; i < n; i++ {
		t := -1
		for j := 0; j < n; j++ {
			if !used[j] && (t == -1 || dis[j] < dis[t]) {
				t = j
			}
		}

		used[t] = true

		for i := 0; i < len(g[t]); i++ {
			dis[i] = min(dis[i], dis[t]+g[t][i])
		}
	}

	for _, di := range dis {
		if di == inf {
			return -1
		}
		ans = max(ans, di)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
