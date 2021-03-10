package main

import "math"

// 1、用到中序遍历，因为中序遍历后是有序的，最小差一定是有序数组两个相邻元素的差。
func getMinimumDifference(root *TreeNode) int {
	res, pre := math.MaxInt64, -1
	dfs(root, &res, &pre)
	return res
}

func dfs(node *TreeNode, res, pre *int) {
	if node == nil { return }

	dfs(node.Left, res, pre)

	// 计算最小值，
	if *pre != -1 && node.Val-*pre < *res {
		*res = node.Val - *pre
	}
	*pre = node.Val

	dfs(node.Right, res, pre)
}
