package main

var e = [100]int{}
var l = [100]int{}
var r = [100]int{}
var index = 0

func init() {
	r[0] = 1
	l[1] = 0

}

func add(k, x int) {
	e[index] = x
	// index 的 右边 是 k 的 右边
	r[index] = r[k]
	// index 的 左边 是 k
	l[index] = k
	// k 右边的左边 是 index
	l[r[k]] = index
	// k 的右边的 index
	r[k] = index
}

func remove(k int) {
	r[l[k]] = r[k]
	l[r[k]] = l[k]
}
