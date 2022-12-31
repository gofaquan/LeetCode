package ___offer

// dfs
func movingCount(m int, n int, k int) int {
	check := make([][]bool, m)
	for i := 0; i < m; i++ {
		check[i] = make([]bool, n)
	}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || check[x][y] == true {
			return 0
		}
		check[x][y] = true
		if (x/10 + x%10 + y/10 + y%10) > k {
			return 0
		}
		return 1 + dfs(x+1, y) + dfs(x-1, y) + dfs(x, y+1) + dfs(x, y-1)
	}

	return dfs(0, 0)
}

// bfs
var dx = []int{1, 0, -1, 0}
var dy = []int{0, 1, 0, -1}

func movingCount2(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var queue [][]int
	queue = append(queue, []int{0, 0})
	visited[0][0] = true

	count := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		count++

		for i := 0; i < 4; i++ {
			x, y := curr[0]+dx[i], curr[1]+dy[i]
			if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] && sum(x)+sum(y) <= k {
				visited[x][y] = true
				queue = append(queue, []int{x, y})
			}
		}
	}

	return count
}

func sum(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}
