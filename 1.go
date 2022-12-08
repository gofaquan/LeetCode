package main

func qs(intervals [][]int, l, r int) {
	if l >= r {
		return
	}
	x, i, j := intervals[l][0], l-1, r+1

	for i < j {
		for i++; intervals[i][0] < x; i++ {
		}
		for j--; intervals[j][0] > x; j-- {
		}

		if i < j {
			intervals[i], intervals[j] = intervals[j], intervals[j]
		}
	}

	qs(intervals, l, j)
	qs(intervals, j+1, r)
}

func main() {

}
