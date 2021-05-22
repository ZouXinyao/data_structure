package main

// 与最大深度不同的是，当遇到只有1个子节点，走到那个为nil的分支上时，最大深度会直接算出来然后舍弃掉。
// 但是最小深度，只有1个分支时不是最低点，需要左右两个都是空才是。所以对1个分支进行了单独处理。
func minDepth(root *TreeNode) int {
	if root == nil { return 0 }
	if root.Left != nil && root.Right == nil { return minDepth(root.Left) + 1 }
	if root.Left == nil && root.Right != nil { return minDepth(root.Right) + 1 }
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(a, b int) int {
	if a < b { return a }
	return b
}