package leecode

import "math"

func networkDelayTime(times [][]int, n, k int) (ans int) {
	const inf = math.MaxInt64 / 2
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				g[i][j] = 0
			} else {
				g[i][j] = inf
			}
		}
	}

	for _, time := range times {
		g[time[0]-1][time[1]-1] = time[2]
	}

	for p := 0; p < n; p++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				g[i][j] = min(g[i][j], g[i][p]+g[p][j])
			}
		}
	}

	for i := 0; i < n; i++ {
		ans = max(g[k-1][i], ans)
	}

	if ans >= inf {
		return -1
	} else {
		return ans
	}
	return
}

type Edge struct {
	a, b, c int
}

func networkDelayTime(times [][]int, n, k int) (ans int) {
	var set []Edge
	const inf = math.MaxInt64 / 2

	for _, time := range times {
		set = append(set, Edge{
			a: time[0] - 1,
			b: time[1] - 1,
			c: time[2],
		})
	}

	dis := make([]int, n)
	for i := 0; i < len(dis); i++ {
		dis[i] = inf
	}
	dis[k-1] = 0

	for i := 0; i < n; i++ {
		var cp = make([]int, n)
		copy(cp, dis)
		for _, edge := range set {
			a, b, c := edge.a, edge.b, edge.c
			dis[b] = min(dis[b], cp[a]+c)
		}

	}

	for _, v := range dis {
		if v >= inf {
			return -1
		}

		ans = max(ans, v)
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
