package main

import (
	"fmt"
	"sort"
)

type PII struct {
	first, second int
}

const N = 300010

var (
	a     [N]int
	s     [N]int
	n, m  int
	alls  []int
	add   []PII
	query []PII
)

func find(x int) int {
	l, r := 0, len(alls)-1
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
	fmt.Scan(&n, &m)

	for i := 0; i < n; i++ {
		var x, c int
		fmt.Scan(&x, &c)
		add = append(add, PII{x, c})

		alls = append(alls, x)
	}

	for i := 0; i < m; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		query = append(query, PII{l, r})

		alls = append(alls, l)
		alls = append(alls, r)
	}

	sort.Ints(alls)
	unique(alls)

	for _, item := range add {
		x := find(item.first)
		a[x] += item.second
	}

	for i := 1; i <= len(alls); i++ {
		s[i] = s[i-1] + a[i]
	}

	for _, item := range query {
		l, r := find(item.first), find(item.second)
		fmt.Println(s[r] - s[l-1])
	}
}

func unique(s []int) {
	var a []int
	m := map[int]int{}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			continue
		}
		m[s[i]]++
		a = append(a, s[i])
	}

	s = a
}
