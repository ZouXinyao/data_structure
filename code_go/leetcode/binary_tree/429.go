package main

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil { return res }
	queue := []*Node{root}
	level := 0
	for len(queue) > 0 {
		res = append(res, []int{})
		length := len(queue)
		for i:=0; i<length; i++ {
			node := queue[0]
			queue = queue[1:]
			res[level] = append(res[level], node.Val)
			for _, chil := range node.Children {
				queue = append(queue, chil)
			}
		}
		level++
	}
	return res
}
