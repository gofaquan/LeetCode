package main

import (
	"fmt"
	"testing"
)

var _94state [9]bool
var _94n int
var _94data [9]int

func _94dfs(u int) {
	if u == _94n {
		for i := 0; i < _94n; i++ {
			fmt.Print(_94data[i], " ")
		}
		fmt.Println()
		return
	}

	for i := 0; i < _94n; i++ {
		if !_94state[i] {
			_94state[i] = true
			_94data[u] = i + 1
			_94dfs(u + 1)
			_94state[i] = false
		}
	}
}

func _94main() {
	fmt.Scanf("%d", &_94n)
	_94dfs(0)
}

func TestDfg(t *testing.T) {
	_94n = 3
	_94dfs(0)
}
