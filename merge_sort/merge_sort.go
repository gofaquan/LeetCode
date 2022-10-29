package main

import "fmt"

var tmp [100000]int

func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)

	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			k, i = k+1, i+1
		} else {
			tmp[k] = nums[j]
			k, j = k+1, j+1
		}
	}

	for i <= mid {
		tmp[k] = nums[i]
		k, i = k+1, i+1
	}

	for j <= right {
		tmp[k] = nums[j]
		k, j = k+1, j+1
	}

	for i, j = left, 0; i <= right; i, j = i+1, j+1 {
		nums[i] = tmp[j]
	}
}
func main() {
	var q = []int{3, 2, 1, 4, 5}
	mergeSort(q, 0, len(q)-1)
	fmt.Println(q)
}
