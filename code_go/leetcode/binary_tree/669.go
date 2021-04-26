package main

// 注意点：当前val<low时，需要保留右子树；当前val>high时，需要保留左子树
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil { return root }
	// 这行才是处理当前层。
	if root.Val > high { return trimBST(root.Left, low, high) }
	if root.Val < low  { return trimBST(root.Right, low, high) }

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}