package main

import (
	"fmt"
	"sort"
)

const N = 300010

var _802a [N]int
var _802s [N]int
var _802n, _802m int
var _802alls []int
var _802add []_802pair
var _802query []_802pair

type _802pair struct {
	x int
	c int
}

func _802find(x int) int {
	l := 0
	r := len(_802alls) - 1
	for l < r {
		mid := (l + r) >> 1
		if _802alls[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r + 1
}

func _802main() {
	fmt.Scanf("%d%d", &_802n, &_802m)
	for i := 0; i < _802n; i++ {
		var x, c int
		fmt.Scanf("%d%d", &x, &c)
		_802add = append(_802add, _802pair{x, c})
		_802alls = append(_802alls, x)
	}
	for i := 0; i < _802m; i++ {
		var l, r int
		fmt.Scanf("%d%d", &l, &r)
		_802query = append(_802query, _802pair{l, r})
		_802alls = append(_802alls, l)
		_802alls = append(_802alls, r)
	}
	sort.Ints(_802alls)
	_802alls = _802unique(_802alls)
	for i := 0; i < len(_802add); i++ {
		x := _802find(_802add[i].x)
		_802a[x] += _802add[i].c
	}
	for i := 1; i <= len(_802alls); i++ {
		_802s[i] = _802s[i-1] + _802a[i]
	}
	for i := 0; i < len(_802query); i++ {
		l := _802find(_802query[i].x)
		r := _802find(_802query[i].c)
		fmt.Println(_802s[r] - _802s[l-1])
	}
}

func _802unique(intSlice []int) []int {
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
