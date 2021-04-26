package main


// 一样使用中序遍历，但是传入的参数有点多，所以使用了将函数写成变量的方式。
func findMode(root *TreeNode) []int {
	var res []int
	var cnt int
	var maxCnt int
	var preNode *TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil { return }

		dfs(node.Left)


		if preNode == nil {
			// preNode还没有值时候
			cnt = 1
		} else if preNode.Val == node.Val {
			// 相等就+1
			cnt++
		} else {
			// 不等就回到1
			cnt = 1
		}
		preNode = node

		if cnt == maxCnt {
			// 当前数和最大数一样，结果中加一个，
			res = append(res, node.Val)
		}
		if cnt > maxCnt {
			// 当前数大，就清空res，把新的结果放进去
			maxCnt = cnt
			res = []int{node.Val}
		}

		dfs(node.Right)
	}

	dfs(root)
	return res
}
