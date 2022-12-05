package main

import (
	"fmt"
	"testing"
)

var _92state int
var _92n int

func _92dfs(u int, state int) {
	if u == _92n {
		for i := 0; i < _92n; i++ {
			if (state >> i & 1) == 1 {
				fmt.Print(i+1, " ")
			}
		}
		fmt.Println()
		return
	}

	_92dfs(u+1, state)
	_92dfs(u+1, state|1<<u)
}

func _92main() {
	fmt.Scanf("%d", &_92n)
	_92dfs(0, 0)
}

func TestBfs(t *testing.T) {
	_92n = 3
	_92dfs(0, 0)
}
