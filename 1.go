package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}
type ListNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
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

func main() {
	postorder := []int{9, 15, 7, 20, 3}
	fmt.Println(verifyPostorder(postorder)) // 输出 true

}
