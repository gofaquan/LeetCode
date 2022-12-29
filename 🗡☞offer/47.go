package ___offer

func maxValue(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	// dp[i][j] represents the maximum value from grid[0][0] to grid[i-1][j-1]
	dp := make([][]int, row+1)
	for i := range dp {
		dp[i] = make([]int, column+1)
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[row][column]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
