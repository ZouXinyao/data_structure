package main

// 递归的模板
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 1、递归终止条件
	if root == nil { return false }
	// 2、当前层的处理，也可以是第二个终止条件，
	if root.Left == nil && root.Right == nil { return targetSum == root.Val }

	// 3、递归到下一层。
	return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}