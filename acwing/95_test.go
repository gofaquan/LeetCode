package main

import (
	"fmt"
	"testing"
)

var _95dx = []int{-1, 0, 1, 0, 0}
var _95dy = []int{0, 1, 0, -1, 0}

var _95g [6][6]byte
var _95backup [6][6]byte

func turn(x, y int) {
	for i := 0; i < 5; i++ {
		newx := x + _95dx[i]
		newy := y + _95dy[i]
		if newx < 0 || newx >= 5 || newy < 0 || newy >= 5 {
			continue
		}

		_95g[newx][newy] ^= 1 // 1 变 0， 0 变 1
	}
}

func Test95(tx *testing.T) {
	var t = 1
	fmt.Scanf("%d", &t)

	for ; t > 0; t-- {
		for i := 0; i < 5; {
			var s string
			fmt.Scanln(&s)
			//fmt.Println(_802s)
			for j := 0; j < len(s); j++ {
				_95g[i][j] = s[j] - '0'
			}
			i++
		}
		fmt.Println(_95g)
		res := 10
		for op := 0; op < 32; op++ {
			copy(_95backup[:], _95g[:])
			step := 0
			for i := 0; i < 5; i++ {
				if op>>i&1 == 1 {
					step++
					turn(0, i)
				}
			}

			for i := 0; i < 4; i++ {
				for j := 0; j < 5; j++ {
					if _95g[i][j] == 0 {
						step++
						turn(i+1, j)
					}
				}
			}

			dark := false
			for i := 0; i < 5; i++ {
				if _95g[4][i] == 0 {
					dark = true
					break
				}
			}

			if !dark {
				res = min(res, step)
			}
			copy(_95g[:], _95backup[:])
		}

		if res > 6 {
			res = -1
		}

		fmt.Println(res)

	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
