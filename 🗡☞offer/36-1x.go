package ___offer

func singleNumbers(nums []int) []int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result ^= nums[i]
	}

	bit := result & (-result)
	var num1, num2 int
	for _, num := range nums {
		if num&bit == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}
	return []int{num1, num2}
}
