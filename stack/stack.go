package main

import "fmt"

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

// 单调栈
var stk = [100]int{}
var input = [5]int{3, 4, 2, 7, 5}
var t = 0

func main() {
	for i := 0; i < 5; i++ {
		for t > 0 && stk[t] >= input[i] {
			t--
		}

		if t > 0 {
			fmt.Println(stk[t], " ")
		} else {
			fmt.Println(-1, " ")
		}

		t++
		stk[t] = input[i]
	}

	fmt.Println(stk)
}
