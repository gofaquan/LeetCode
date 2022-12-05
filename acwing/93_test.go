package main

import (
	"fmt"
	"testing"
)

var _93data [26]int
var _93n int
var _93m int

func _93dfs(u, start int) {
	if u+_93n-start < _93m {
		return
	}

	if u == _93m {
		for i := 0; i < _93m; i++ {
			fmt.Print(_93data[i], " ")
		}
		fmt.Println()
		return
	}

	for i := start; i < _93n; i++ {
		_93data[u] = i + 1
		_93dfs(u+1, i+1)
		_93data[u] = 0
	}
}

func main() {
	fmt.Scan(&_93n)
	fmt.Scan(&_93m)
	_93dfs(0, 0)
}

func Test93(t *testing.T) {
	//fmt.Scan(&_93n)
	//fmt.Scan(&_93m)
	_93n = 5
	_93m = 3

	_93dfs(0, 0)
}
