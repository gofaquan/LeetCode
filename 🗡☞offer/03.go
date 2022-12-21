package ___offer

func findRepeatNumber(nums []int) (res int) {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
			continue
		}

		if nums[nums[i]] == nums[i] {
			return nums[i]
		}

		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}

	return -1
}
