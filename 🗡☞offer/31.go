package ___offer

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	j := 0
	for _, x := range pushed {
		stack = append(stack, x)
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}

	return len(stack) == 0
}
