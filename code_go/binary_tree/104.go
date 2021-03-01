package main

// 理解求最大深度就是递归的求当前节点max(左子树最大深度，右子树最大深度)+1
func maxDepth(root *TreeNode) int {
	if root == nil { return 0 }

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b { return a }
	return b
}