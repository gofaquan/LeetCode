package main

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func firstUniqChar(s string) byte {
	m := map[uint8]int{}
	for i, j := 0, 0; j < len(s); {
		if m[s[j]] == 1 {
			m[s[i]]--
		}
	}
}

func main() {
	var a [][]int
	fmt.Println(len(a[0]))
}
