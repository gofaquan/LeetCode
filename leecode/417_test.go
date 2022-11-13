package leecode

var direction = []struct{ x, y int }{
	{-1, 0}, //左
	{1, 0},  // 右
	{0, -1}, // 下
	{0, 1},  // 上
}

// 深度优先搜索 dfs
func pacificAtlantic(heights [][]int) (ans [][]int) {
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func(x int, y int, ocean [][]bool)
	dfs = func(x int, y int, ocean [][]bool) {
		if ocean[x][y] {
			return
		}
		ocean[x][y] = true
		for _, d := range direction {
			if nx, ny := x+d.x, y+d.y; 0 <= nx && nx < m && 0 <= ny && ny < n &&
				heights[nx][ny] >= heights[x][y] {
				dfs(nx, ny, ocean)
			}
		}
	}

	// 左边的 pacific
	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
	}
	// 上边的 pacific , [0,0] 必是
	for j := 1; j < n; j++ {
		dfs(0, j, pacific)
	}
	// 右边的 atlantic
	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic)
	}
	// 下边的 atlantic , [0,n-1] 必是
	for j := 0; j < n-1; j++ {
		dfs(m-1, j, atlantic)
	}

	for i, row := range pacific {
		for j, ok := range row {
			if ok && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}
