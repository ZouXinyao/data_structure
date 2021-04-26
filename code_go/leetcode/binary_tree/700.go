package main

// 一个简单的递归
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || val == root.Val {
		return root
	} else if val < root.Val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
