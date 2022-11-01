package main

var q = [100]int{}
var head, end int

func enQueue(x int) {
	head++
	q[head] = x
}

func deQueue() int {
	x := q[head]
	head--
	return x
}

func isEmpty() bool {
	if head > end {
		return false
	}
	return true
}
func getHead() int {
	return q[head]
}
