package main

import (
	"fmt"
	"testing"
)

var _717data []int
var _717n int

func fibonacci(n, now int, data []int) {
	if now == n {
		return
	}

	if now == 0 {
		fmt.Print(0, " ")
		data[now] = 0
	} else if now == 1 {
		fmt.Print(1, " ")
		data[now] = 1
	} else {
		data[now] = data[now-1] + data[now-2]
		fmt.Print(data[now], " ")
	}

	fibonacci(n, now+1, data)
}

func main() {
	fmt.Scanf("%d", &_717n)
	_717data = make([]int, _717n)

	fibonacci(_717n, 0, _717data)
}

func Test717(t *testing.T) {
	_717n = 5
	_717data = make([]int, _717n)
	fibonacci(_717n, 0, _717data)
}
