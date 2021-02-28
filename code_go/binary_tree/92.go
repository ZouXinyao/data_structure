package main

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil { return res }
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}