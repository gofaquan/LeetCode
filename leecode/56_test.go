package leecode

import (
	"fmt"
	"testing"
)

var testArray = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

func TestQuickSort(t *testing.T) {
	_56quickSort(testArray, 0, len(testArray)-1)
	fmt.Println(testArray)
}

func TestMerge(t *testing.T) {
	//fmt.Println(merge(testArray))
	fmt.Println(_56merge2([][]int{{1, 4}, {0, 4}}))
}

func _56merge2(intervals [][]int) (res [][]int) {
	if len(intervals) == 1 {
		return intervals
	}

	_56quickSort(intervals, 0, len(intervals)-1)

	left := intervals[0][0]
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right {
			if intervals[i][1] > right {
				right = intervals[i][1]
			}
		} else {
			res = append(res, []int{left, right})
			right = intervals[i][1]
			left = intervals[i][0]
		}
	}
	res = append(res, []int{left, right})
	return
}

func _56merge1(intervals [][]int) [][]int {
	var res [][]int
	_56quickSort(intervals, 0, len(intervals)-1)
	// 0 <= starti <= endi <= 10^4
	st, end := -1, -1

	for _, interval := range intervals {
		if end < interval[0] {
			if st != -1 {
				res = append(res, []int{st, end})
			}
			st, end = interval[0], interval[1]
		} else {
			end = _57max(interval[1], end)
		}
	}

	if st != -1 {
		res = append(res, []int{st, end})
	}

	return res
}

func _56quickSort(intervals [][]int, left, right int) {
	if left >= right {
		return
	}

	i, j, x := left-1, right+1, intervals[left][0]

	for i < j {
		for i++; intervals[i][0] < x; i++ {
		}
		for j--; intervals[j][0] > x; j-- {
		}

		if i < j {
			intervals[i], intervals[j] = intervals[j], intervals[i]
		}
	}

	_56quickSort(intervals, left, j)
	_56quickSort(intervals, j+1, right)
}
