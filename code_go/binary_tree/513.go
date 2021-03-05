package main

// BFS
func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{}
	if root != nil { queue = append(queue, root) }
	result := 0
	for len(queue) > 0 {
		size := len(queue)
		for i:=0; i<size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == 0 { result = node.Val }
			if node.Left != nil { queue = append(queue, node.Left) }
			if node.Right != nil {  queue = append(queue, node.Right) }
		}
	}
	return result
}
