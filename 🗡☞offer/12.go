package ___offer

// 优化空间
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var dfs func(x, y, next int) bool
	dfs = func(x, y, next int) bool {
		if next == len(word) {
			return true
		}
		if x < 0 || x == len(board) || y < 0 ||
			y == len(board[0]) || board[x][y] != word[next] {
			return false
		}
		tmp := board[x][y]
		board[x][y] = 0
		res := dfs(x+1, y, next+1) || dfs(x-1, y, next+1) || dfs(x, y+1, next+1) || dfs(x, y-1, next+1)
		board[x][y] = tmp
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// 多了空间
func exist2(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	check := make([][]bool, m)
	for i := 0; i < m; i++ {
		check[i] = make([]bool, n)
	}

	var dfs func(x, y, next int) bool
	dfs = func(x, y, next int) bool {
		if next == len(word) {
			return true
		}
		if x < 0 || x == len(board) || y < 0 || y == len(board[0]) || board[x][y] != word[next] || check[x][y] == true {
			return false
		}
		check[x][y] = true
		res := dfs(x+1, y, next+1) || dfs(x-1, y, next+1) || dfs(x, y+1, next+1) || dfs(x, y-1, next+1)
		check[x][y] = false
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
