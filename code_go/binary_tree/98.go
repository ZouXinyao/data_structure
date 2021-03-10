package main

import "math"

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)

}

// 递归的思路，
func helper(node *TreeNode, min, max int) bool {
	// 1、两个终止条件，如果遍历结束，就返回true，说明遍历的部分都满足
	// 2、当遇到不满足条件的直接返回false
	if node == nil { return true }
	if node.Val <= min || node.Val >= max { return false }

	// 结果就是左子树满足条件&&右子树满足条件，计算左子树就是有最大值了，计算右子树就是有最小值了。
	return helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
}
