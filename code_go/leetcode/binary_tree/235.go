package main

// 二叉搜索树的性质，左子树都小于父，右子树都大于父。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res := root
	for {
		if res.Val > q.Val && res.Val > p.Val {
			res = res.Left
		} else if res.Val < q.Val && res.Val < p.Val {
			res = res.Right
		} else { return res }
	}
}
