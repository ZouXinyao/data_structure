package main

func isSymmetric(root *TreeNode) bool {
	if root == nil { return true }
	return check(root.Left, root.Right)
}

func check(l *TreeNode, r *TreeNode) bool {
	// 是否对称的几种情况，
	if l == nil && r == nil { return true }
	if l == nil && r != nil { return false }
	if l != nil && r == nil { return false }

	if l.Val != r.Val { return false }
	// 需要递归的是从中间看是否对称，而不是当前节点的左右对称
	return check(l.Left, r.Right) && check(l.Right, r.Left)
}