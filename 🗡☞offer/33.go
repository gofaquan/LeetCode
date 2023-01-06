package ___offer

func verifyPostorder(postorder []int) bool {
	// 如果数组为空，则认为是后序遍历结果
	if len(postorder) == 0 {
		return true
	}

	// 取出数组的最后一个元素，作为当前二叉搜索树的根节点
	root := postorder[len(postorder)-1]

	// 从前往后搜索，找到第一个比根节点小的元素
	i := 0
	for i < len(postorder)-1 && postorder[i] < root {
		i++
	}

	// 从这个元素的下一个位置开始，找到第一个比根节点大的元素
	for j := i; j < len(postorder)-1; j++ {
		// 如果在搜索过程中找到了比根节点小的元素，则说明数组不是后序遍历结果
		if postorder[j] < root {
			return false
		}
	}

	// 递归地判断根节点的左子树是否符合要求
	leftIsValid := true
	if i > 0 {
		leftIsValid = verifyPostorder(postorder[:i])
	}

	// 递归地判断根节点的右子树是否符合要求
	rightIsValid := true
	if i < len(postorder)-1 {
		rightIsValid = verifyPostorder(postorder[i : len(postorder)-1])
	}

	// 如果左子树和右子树都符合要求，则说明整棵树的后序遍历结果是正确的
	return leftIsValid && rightIsValid
}
