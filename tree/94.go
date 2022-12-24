package tree

type Pair struct {
	node  *TreeNode
	color int
}

func inorderTraversal(root *TreeNode) (res []int) {
	const white, gray = 0, 1
	stack := []Pair{{root, white}}
	for len(stack) > 0 {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pair.node == nil {
			continue
		}
		if pair.color == white {
			stack = append(stack, Pair{node: pair.node.Right, color: white})
			stack = append(stack, Pair{node: pair.node, color: gray})
			stack = append(stack, Pair{node: pair.node.Left, color: white})
		} else {
			res = append(res, pair.node.Val)
		}

	}

	return res
}
