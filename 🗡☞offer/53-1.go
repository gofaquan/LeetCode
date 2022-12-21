package ___offer

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	x := bsearch1(nums, target)
	if nums[x] != target {
		return 0
	}
	y := bsearch2(nums, target)
	return y - x + 1
}
func bsearch1(nums []int, x int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] < x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func bsearch2(nums []int, x int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right + 1) / 2
		if nums[mid] <= x {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	println(search(nums, 8))
}
