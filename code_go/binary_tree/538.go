package main

// 反向中序遍历，新二叉树每个节点的val是比他大节点的和。反序求和：
// 二叉搜索树想到中序遍历，因为是有序的。
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil { return }
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}


