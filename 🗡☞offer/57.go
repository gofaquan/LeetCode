package ___offer

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		res := nums[l] + nums[r]

		if res > target {
			r--
		} else if res < target {
			l++
		} else {
			return []int{nums[l] + nums[r]}
		}
	}

	return nil
}
