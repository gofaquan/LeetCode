package ___offer

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	isReverse := false

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

		if isReverse {
			for i := 0; i < len(level)/2; i++ {
				level[i], level[len(level)-i-1] = level[len(level)-i-1], level[i]
			}
		}

		isReverse = !isReverse
		res = append(res, level)
	}

	return res
}
