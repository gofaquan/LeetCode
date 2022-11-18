package leecode

import (
	"fmt"
	"testing"
)

var testArray = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

func TestQuickSort(t *testing.T) {
	quickSort(testArray, 0, len(testArray)-1)
	fmt.Println(testArray)
}

func TestMerge(t *testing.T) {
	fmt.Println(merge(testArray))
}

func merge(intervals [][]int) [][]int {
	var res [][]int
	quickSort(intervals, 0, len(intervals)-1)
	// 0 <= starti <= endi <= 10^4
	st, end := -1, -1

	for _, interval := range intervals {
		if end < interval[0] {
			if st != -1 {
				res = append(res, []int{st, end})
			}
			st, end = interval[0], interval[1]
		} else {
			end = max(interval[1], end)
		}
	}

	if st != -1 {
		res = append(res, []int{st, end})
	}

	return res
}

var tmp [100000]int

func quickSort(queue [][]int, left, right int) {
	if left >= right {
		return
	}

	leftBoard, rightBoard, Val := left-1, right+1, queue[left][0]

	for leftBoard < rightBoard {
		for leftBoard++; queue[leftBoard][0] < Val; leftBoard++ {
		}
		for rightBoard--; queue[rightBoard][0] > Val; rightBoard-- {
		}

		if leftBoard < rightBoard {
			queue[leftBoard], queue[rightBoard] = queue[rightBoard], queue[leftBoard]
		}
	}

	quickSort(queue, left, rightBoard)
	quickSort(queue, rightBoard+1, right)
}
