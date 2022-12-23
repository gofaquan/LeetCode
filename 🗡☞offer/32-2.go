package ___offer

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		var level []int
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}

	return res
}
