package leecode

import (
	"fmt"
	"testing"
)

type _303NumArray struct {
	sums []int
}

func _303Constructor(nums []int) (res _303NumArray) {
	res.sums = make([]int, len(nums))
	res.sums[0] = nums[0]

	for i := 0; i < len(nums)-1; i++ {
		res.sums[i+1] = res.sums[i] + nums[i+1]
	}

	return
}

func (na *_303NumArray) SumRange(i, j int) int {
	if i == 0 {
		return na.sums[j]
	}

	return na.sums[j] - na.sums[i-1]
}

func TestSumRange(t *testing.T) {
	a := []int{1, -1, 3}
	res := _303Constructor(a)
	fmt.Println(res.SumRange(0, 1))
	fmt.Println(res.sums)
}
