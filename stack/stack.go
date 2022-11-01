package main

var stack = [100]int{}
var top int

func push(x int) {
	top++
	stack[top] = x
}

func isEmpty() bool {
	if top > 0 {
		return false
	}
	return true
}
func getTopVal() int {
	return stack[top]
}

func pop() int {
	x := stack[top]
	top--
	return x
}
