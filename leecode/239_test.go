package leecode

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	var queue = make([]int, len(nums))
	head, tail := 0, -1

	for i := 0; i < len(nums); i++ {
		if head <= tail && i-k+1 > queue[head] {
			head++
		}

		for head <= tail && nums[queue[tail]] <= nums[i] {
			tail--
		}
		tail++
		queue[tail] = i

		if i >= k-1 {
			result = append(result, nums[queue[head]])
		}
	}

	return result
}

/*
if hh <= tt
*/
