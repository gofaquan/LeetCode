package tree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)

	return result
}
