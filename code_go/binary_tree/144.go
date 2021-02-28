package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 三种递归写法

func preorderTraversal1(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal1(root.Left)...)
	res = append(res, preorderTraversal1(root.Right)...)
	return
}

var res []int
func preorderTraversal2(root *TreeNode) []int {
	res = []int{}
	dfs(root)
	return res
}
func dfs(root *TreeNode) {
	if root != nil {
		res = append(res, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
}

func preorderTraversal3(root *TreeNode) (res []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

