package ___offer

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {

		for left < right && nums[left]%2 == 1 {
			left++
		}

		for left < right && nums[right]%2 == 0 {
			right--
		}

		if nums[left]%2 == 0 && nums[right]%2 == 1 {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	return nums
}
