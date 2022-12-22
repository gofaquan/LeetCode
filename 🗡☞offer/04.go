package ___offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}

	for i := 0; i < n; i++ {
		if matrix[i][0] <= target && matrix[i][m-1] >= target {

			left, right := 0, m-1
			for left <= right {
				mid := (left + right) >> 1
				if matrix[i][mid] == target {
					left = mid
					break
				} else if matrix[i][mid] > target {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}

			if matrix[i][left] == target {
				return true
			}

		}

		if matrix[i][0] > target {
			return false
		}
	}
	return false
}
