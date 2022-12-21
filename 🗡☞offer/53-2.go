package ___offer

import "fmt"

func missingNumber(nums []int) (res int) {

	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}

func main() {
	fmt.Println(missingNumber([]int{1}))
}
