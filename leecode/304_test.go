package leecode

import (
	"fmt"
	"testing"
)

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sum := make([][]int, len(matrix))
	for i := range sum {
		sum[i] = make([]int, len(matrix[0]))
	}

	sum[0][0] = matrix[0][0]
	for i := 1; i < len(matrix); i++ {
		sum[i][0] = sum[i-1][0] + matrix[i][0]
	}
	for i := 1; i < len(matrix[0]); i++ {
		sum[0][i] = sum[0][i-1] + matrix[0][i]
	}

	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[0]); col++ {
			sum[row][col] = sum[row-1][col] + sum[row][col-1] + matrix[row][col] - sum[row-1][col-1]
		}
	}

	return NumMatrix{sums: sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	s := this.sums
	if row1 == 0 && col1 == 0 {
		return s[row2][col2]
	}

	if row1 == 0 {
		return s[row2][col2] - s[row2][col1-1]
	}
	if col1 == 0 {
		return s[row2][col2] - s[row1-1][col2]
	}

	return s[row2][col2] - s[row1-1][col2] - s[row2][col1-1] + s[row1-1][col1-1]
}

func TestNumm(t *testing.T) {
	a := [][]int{{-4, -5}}
	constructor := Constructor(a)
	fmt.Println(constructor.SumRegion(0, 1, 0, 1))
}
