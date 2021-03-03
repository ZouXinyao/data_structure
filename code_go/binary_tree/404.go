package main

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	dfs(root, &res, false)
	return res
}
// 增加1个标志位表示左子树，然后用左右节点为空筛选叶子
func dfs(node *TreeNode, res *int, flag bool) {
	if node == nil { return }
	if flag && (node.Left == nil && node.Right == nil) { *res += node.Val }

	dfs(node.Left, res, true)
	dfs(node.Right, res, false)
}
