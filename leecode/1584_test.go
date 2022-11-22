package leecode

import (
	"fmt"
	"math"
	"sort"
)

type edge struct {
	a, b, c int
}

// failed
func minCostConnectPointsf(points [][]int) (res int) {
	const inf = math.MaxInt64 / 2
	edges := make([]edge, len(points))
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			x1 := points[i][0]
			y1 := points[i][1]
			x2 := points[j][0]
			y2 := points[j][1]
			w := int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
			edges[i] = edge{
				a: i,
				b: j,
				c: w,
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].c < edges[j].c
	})

	fmt.Println(edges)

	p := make([]int, len(points))
	for i := 0; i < len(p); i++ {
		p[i] = i
	}

	var f func(int) int
	f = func(x int) int {
		if p[x] != x {
			p[x] = f(p[x])
		}

		return p[x]
	}

	for _, edge := range edges {
		a, b, c := edge.a, edge.b, edge.c
		a = f(a)
		b = f(b)
		if a != b {
			res += c
			p[a] = b
		}
	}

	return res
}

func minCostConnectPoints(points [][]int) (res int) {
	const inf = math.MaxInt64 / 2
	g := make([][]int, len(points))
	for i := 0; i < len(points); i++ {
		g[i] = make([]int, len(points))
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			x1 := points[i][0]
			x2 := points[j][0]
			y1 := points[i][1]
			y2 := points[j][1]
			w := int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
			g[i][j] = w
			g[j][i] = w
		}
	}

	dis := make([]int, len(points))
	for i := 0; i < len(dis); i++ {
		dis[i] = inf
	}
	used := make([]bool, len(points))

	for i := 0; i < len(points); i++ {
		t := -1
		for j := 0; j < len(points); j++ {
			if !used[j] && (t == -1 || dis[t] > dis[j]) {
				t = j
			}
		}

		if i > 0 {
			res += dis[t]
		}

		for j := 0; j < len(points); j++ {
			dis[j] = min(dis[j], g[t][j])
		}

		used[t] = true
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
