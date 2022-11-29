package leecode

import "testing"

func findGCD(nums []int) int {
	a := sMax(nums)
	b := sMin(nums)

	return gcd(a, b)
}

func gcd(x, y int) int {
	if y > 0 {
		return gcd(y, x%y)
	} else {
		return x
	}

}

func sMax(nums []int) (res int) {
	for _, num := range nums {
		if res < num {
			res = num
		}
	}
	return
}

func sMin(nums []int) (res int) {
	res = nums[0]
	for _, num := range nums {
		if res > num {
			res = num
		}
	}
	return
}

func TestFGcd(t *testing.T) {
	println(findGCD([]int{2, 5, 6, 9, 10}))
}
