package main

import "fmt"

//var q = [100]int{}
//var head, end int
//
//func enQueue(x int) {
//	head++
//	q[head] = x
//}
//
//func deQueue() int {
//	x := q[head]
//	head--
//	return x
//}
//
//func isEmpty() bool {
//	if head > end {
//		return false
//	}
//	return true
//}
//func getHead() int {
//	return q[head]
//}

// 滑动窗口
var a = [8]int{1, 3, -1, -3, 5, 3, 6, 7}
var queue = [100]int{}

func window() {
	hh, tt, k := 0, -1, 3

	for i := 0; i < 8; i++ {
		if hh <= tt && i-k+1 > queue[hh] {
			hh++
		}

		for hh <= tt && a[queue[tt]] >= a[i] {
			tt--
		}

		tt++
		queue[tt] = i

		if i >= k-1 {
			fmt.Printf("%d ", a[queue[hh]])
		}
	}
}

func main() {
	window()
}
