package main

func postorder(root *Node) []int {
	var res []int
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil { return }

		for _, chil := range node.Children {
			dfs(chil)
		}
		res = append(res, node.Val)
	}
	dfs(root)

	return res
}
