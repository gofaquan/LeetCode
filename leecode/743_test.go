package leecode

import "math"

func networkDelayTime(times [][]int, n int, k int) (ans int) {
	const inf = math.MaxInt64 / 2
	d := make([]int, n)
	for i := 0; i < len(d); i++ {
		d[i] = inf
	}
	d[k-1] = 0
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
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
			if !used[j] && (t == -1 || d[t] > d[j]) {
				t = j
			}
		}

		used[t] = true

		for a := 0; a < n; a++ {
			d[a] = min(d[a], d[t]+g[t][a])
		}
	}

	for _, v := range d {
		if v == inf {
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
