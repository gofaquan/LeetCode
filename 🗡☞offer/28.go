package ___offer

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var equal func(l, r *TreeNode) bool
	equal = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}

		if l == nil || r == nil || l.Val != r.Val {
			return false
		}

		return equal(l.Left, r.Right) && equal(l.Right, r.Left)
	}

	return equal(root.Left, root.Right)

}
