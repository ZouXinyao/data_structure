package main

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	buildPath(root, "", &res)
	return res
}

// 标准的递归写法：
func buildPath(node *TreeNode, path string, res *[]string) {
	// 1、递归终止条件
	if node == nil { return }
	if node.Left == nil && node.Right == nil {
		path += strconv.Itoa(node.Val)
		*res = append(*res, path)
		return
	}
	// 2、处理当前层
	path += strconv.Itoa(node.Val) + "->"
	// 3、下探到下一层
	buildPath(node.Left, path, res)
	buildPath(node.Right, path, res)
}
