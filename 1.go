package main

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if abs(maxDepth(root.Left)-maxDepth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
