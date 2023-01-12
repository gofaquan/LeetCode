package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//	type Node struct {
//		Val   int
//		Left  *Node
//		Right *Node
//	}
type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func strToInt(s string) int {
	str := strings.TrimSpace(s)
	if len(str) == 0 {
		return 0
	}
	i := 0

	if !(str[i] == '-' || str[i] == '+' || str[i] >= '0' && str[i] <= '9') {
		return 0
	}
	i++

	for ; i < len(str); i++ {
		if !(str[i] >= '0' && str[i] <= '9') {
			break
		}
	}
	maxStr := strconv.Itoa(math.MaxInt32)
	var newStr string
	if str[0] == '-' {
		newStr = str[1:i]
		if newStr > maxStr && len(newStr) >= len(maxStr) {
			return math.MinInt32
		} else {
			res, _ := strconv.Atoi(newStr)
			return -res
		}
	} else {
		newStr = str[:i]
		if newStr > maxStr || len(newStr) >= len(maxStr) {
			return math.MaxInt32
		} else {
			res, _ := strconv.Atoi(newStr)
			return res
		}
	}

}

func main() {
	fmt.Println(strToInt(" "))
}
