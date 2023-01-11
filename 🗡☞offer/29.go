package ___offer

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return nil
	}
	var ans []int
	cols := len(matrix[0])
	start := 0
	for rows > start*2 && cols > start*2 {
		ans = append(ans, printMatrixInCircle(matrix, rows, cols, start)...)
		start++
	}

	return ans
}

func printMatrixInCircle(matrix [][]int, rows, cols, start int) []int {
	var res []int
	endX := cols - 1 - start
	endY := rows - 1 - start
	// 从左到右打印一行
	for i := start; i <= endX; i++ {
		res = append(res, matrix[start][i])
	}
	// 从上到下打印一列
	if start < endY {
		for i := start + 1; i <= endY; i++ {
			res = append(res, matrix[i][endX])
		}
	}
	// 从右到左打印一行
	if start < endX && start < endY {
		for i := endX - 1; i >= start; i-- {
			res = append(res, matrix[endY][i])
		}
	}
	// 从下到上打印一列
	if start < endX && start < endY-1 {
		for i := endY - 1; i > start; i-- {
			res = append(res, matrix[i][start])
		}
	}
	return res
}
