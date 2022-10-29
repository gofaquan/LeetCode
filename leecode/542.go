package leecode

type point struct {
	x, y int
}

var dirs = []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func updateMatrix(mat [][]int) [][]int {
	var m, n = len(mat), len(mat[0])
	var res = make([][]int, m)
	var visited = make([][]bool, m)
	var queue []point
	for i := range mat {
		res[i] = make([]int, n)
		visited[i] = make([]bool, n)
		for j := range mat[i] {
			if mat[i][j] == 0 {
				queue = append(queue, point{i, j})
				visited[i][j] = true
			}
		}
	}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, dir := range dirs {
			i, j := cur.x+dir.x, cur.y+dir.y
			if i >= 0 && i < m && j >= 0 && j < n && !visited[i][j] {
				res[i][j] = res[cur.x][cur.y] + 1
				queue = append(queue, point{i, j})
				visited[i][j] = true
			}
		}
	}
	return res
}
