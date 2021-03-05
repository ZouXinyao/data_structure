package main

func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

// 所有节点的左右节点互换，直接递归模板。
func invert(node *TreeNode) {
	if node == nil { return }
	node.Left, node.Right = node.Right, node.Left
	invert(node.Left)
	invert(node.Right)
}

