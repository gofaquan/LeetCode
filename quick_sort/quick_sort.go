package main

import "fmt"
func quickSort(queue []int, left, right int) {
	if left >= right {
		return
	}

	x, i, j := queue[left], left-1, right+1
	for i < j {
		for i++; queue[i] < x; i++ {
		}
		for j--; queue[j] > x; j-- {
		}

		if i < j {
			queue[i], queue[j] = queue[j], queue[i]
		}
	}
	quickSort(queue, left, j)
	quickSort(queue, j+1, right)
}

func main() {
	var q = []int{3, 2, 1, 4, 5}
	quickSort(q, 0, len(q)-1)
	fmt.Println(q)
}



