package ___offer

func kthLargest(root *TreeNode, k int) int {
	var result int
	var count int
	var inorder func(node *TreeNode, k int)
	inorder = func(node *TreeNode, k int) {
		if node == nil {
			return
		}

		inorder(node.Right, k)

		count++
		if count == k {
			result = node.Val
			return
		}

		inorder(node.Left, k)
	}

	inorder(root, k)

	return result
}
