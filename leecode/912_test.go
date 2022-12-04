package leecode

import (
	"fmt"
	"testing"
)

func _912sortArray(nums []int) []int {
	//_912QuickSort(nums, 0, len(nums)-1)
	var tmpArr = make([]int, len(nums))

	var MergeSort func(nums []int, l, r int)
	MergeSort = func(nums []int, l, r int) {
		if l >= r {
			return
		}
		mid := (l + r) / 2
		MergeSort(nums, l, mid)
		MergeSort(nums, mid+1, r)

		i, j, k := l, mid+1, 0
		for i <= mid && j <= r {
			if nums[i] > nums[j] {
				tmpArr[k] = nums[j]
				k, j = k+1, j+1
			} else {
				tmpArr[k] = nums[i]
				k, i = k+1, i+1
			}
		}

		for i <= mid {
			tmpArr[k] = nums[i]
			k, i = k+1, i+1
		}
		for j <= r {
			tmpArr[k] = nums[j]
			k, j = k+1, j+1
		}

		for i, j = l, 0; i <= r; i, j = i+1, j+1 {
			nums[i] = tmpArr[j]
		}

	}

	MergeSort(nums, 0, len(nums)-1)
	return nums
}

func _912QuickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	value, i, j := nums[l], l-1, r+1
	for i < j {
		for i++; nums[i] < value; i++ {
		}
		for j--; nums[j] > value; j-- {
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	_912QuickSort(nums, l, j)
	_912QuickSort(nums, j+1, r)
}

var a = []int{5, 2, 3, 1}

func TestQuickSort1(t *testing.T) {
	fmt.Println(_912sortArray(a))
}
