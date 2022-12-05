package main

import (
	"fmt"
	"testing"
)

var state int
var n int

func dfs(u int, state int) {
	if u == n {
		for i := 0; i < n; i++ {
			if (state >> i & 1) == 1 {
				fmt.Print(i+1, " ")
			}
		}
		fmt.Println()
		return
	}

	dfs(u+1, state)
	dfs(u+1, state|1<<u)
}

func main() {
	fmt.Scanf("%d", &n)
	dfs(0, 0)
}

func TestBfs(t *testing.T) {
	n = 3
	dfs(0, 0)
}
