package main

import "fmt"
var tmp [100]int

func mergeSort(queue []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >>1
	mergeSort(queue, left, mid)
	mergeSort(queue, mid+1, right)

	i, j, k := left, mid+1, 0
	for ;i <= mid && j <= right; {
		if queue[i] <= queue[j] {
			tmp[k] = queue[i]
			k++
			i++
		} else {
			tmp[k] = queue[j]
			k++
			j++
		}
	}

	for ;i <= mid; {
		tmp[k] = queue[i]
		k++
		i++
	}
	for ;j <= right; {
		tmp[k] = queue[j]
		k++
		j++
	}

	for i, j = left, 0; i <= right; {
		queue[i] = tmp[j]
		i++
		j++
	}
}

func main() {
	var q = []int{3, 2, 1, 4, 5}
	mergeSort(q, 0, len(q)-1)
	fmt.Println(q)
}
