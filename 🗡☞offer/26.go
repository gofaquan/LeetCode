package ___offer

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var recur func(A *TreeNode, B *TreeNode) bool
	recur = func(A *TreeNode, B *TreeNode) bool {
		if B == nil {
			return true
		}
		if A == nil || A.Val != B.Val {
			return false
		}
		return recur(A.Left, B.Left) && recur(A.Right, B.Right)
	}

	if A == nil || B == nil {
		return false
	}
	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
