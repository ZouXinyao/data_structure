package main

// 由于是完全二叉树，直接数就行；另一种方法是二分查找
func countNodes(root *TreeNode) int {
	if root==nil { return 0 }
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

