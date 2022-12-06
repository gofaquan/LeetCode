package main

import (
	"fmt"
	"sort"
)

const N = 300010

var a [N]int
var s [N]int
var n, m int
var alls []int
var add []pair
var query []pair

type pair struct {
	x int
	c int
}

func find(x int) int {
	l := 0
	r := len(alls) - 1
	for l < r {
		mid := (l + r) >> 1
		if alls[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r + 1
}

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 0; i < n; i++ {
		var x, c int
		fmt.Scanf("%d%d", &x, &c)
		add = append(add, pair{x, c})
		alls = append(alls, x)
	}
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Scanf("%d%d", &l, &r)
		query = append(query, pair{l, r})
		alls = append(alls, l)
		alls = append(alls, r)
	}
	sort.Ints(alls)
	alls = unique(alls)
	for i := 0; i < len(add); i++ {
		x := find(add[i].x)
		a[x] += add[i].c
	}
	for i := 1; i <= len(alls); i++ {
		s[i] = s[i-1] + a[i]
	}
	for i := 0; i < len(query); i++ {
		l := find(query[i].x)
		r := find(query[i].c)
		fmt.Println(s[r] - s[l-1])
	}
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
