package main

// 层序遍历就是BFS，注意保存每层queue的长度，因为queue是变化的，所以不能在for中使用
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil { return res }
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		res = append(res, []int{})
		length := len(queue)
		for i:=0; i<length; i++ {
			node := queue[0]
			queue = queue[1:]
			res[level] = append(res[level], node.Val)
			if node.Left  != nil { queue = append(queue, node.Left) }
			if node.Right != nil { queue = append(queue, node.Right) }
		}
		level++
	}
	return res
}