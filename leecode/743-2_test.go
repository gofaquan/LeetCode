package leecode

import "math"

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
