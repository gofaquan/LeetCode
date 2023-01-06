package ___offer

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	if len(preorder) == 1 {
		return root
	}
	rootIndex := -1
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}
	if rootIndex == -1 {
		return nil
	}
	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]
	leftPreorder := preorder[1 : len(leftInorder)+1]
	rightPreorder := preorder[len(leftInorder)+1:]
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return root
}

// 迭代
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}
	inorderIndex := 0
	for i := 1; i < len(preorder); i++ {
		node := &TreeNode{Val: preorder[i]}
		last := stack[len(stack)-1]
		if last.Val != inorder[inorderIndex] {
			last.Left = node
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				last = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			last.Right = node
		}
		stack = append(stack, node)
	}
	return root
}
