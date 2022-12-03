package leecode

import (
	"fmt"
	"testing"
)

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
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

	quickSort(nums, l, j)
	quickSort(nums, j+1, r)

}

func TestSortArray(t *testing.T) {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
}
