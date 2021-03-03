package main

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

// 所谓的自底向上递归。
// 但是理解成，只需要求出左子树高度和右子树高度，保证差的绝对值<=1，就返回高度。
func height(node *TreeNode) int {
	if node == nil { return 0 }

	leftHeight := height(node.Left)
	rightHeight := height(node.Right)

	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight - rightHeight) > 1 { return -1 }
	return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b { return a}
	return b
}

func abs(x int) int {
	if x < 0 { return -x }
	return x
}
