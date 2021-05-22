package main

// 自底向上传结果。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 终止条件有两个，1个是空节点时，2是当root走到pq时
	if root == nil { return nil }
	if root.Val == p.Val || root.Val == q.Val { return root }

	// 直接下探到下一层，得到在左子树和右子树中最近公共祖先
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	// 这里是求这个公共祖先，要联想到上面的终止条件：
	// 1、相等也终止，表示lr就是pq，root就是结果
	// 2、一边为空表示，另一边一定是结果。
	if l != nil && r != nil { return root }
	if l == nil && r != nil { return r }
	return l
}
