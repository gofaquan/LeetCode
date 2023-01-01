package ___offer

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	var val []int
	var dfs func(a *TreeNode, val []int, sum int)
	dfs = func(a *TreeNode, val []int, sum int) {
		val = append(val, a.Val)
		sum += a.Val
		if a.Left == nil && a.Right == nil {
			if sum == target {
				// 关键cp
				data := make([]int, len(val))
				copy(data, val)
				ans = append(ans, data)
			}
			return
		}

		if a.Left != nil {
			dfs(a.Left, val, sum)
		}

		if a.Right != nil {
			dfs(a.Right, val, sum)
		}

	}

	dfs(root, val, 0)
	return ans
}
