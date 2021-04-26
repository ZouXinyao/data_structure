package main

func preorder(root *Node) []int {
	var res []int
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil { return }
		res = append(res, node.Val)

		for _, chil := range node.Children {
			dfs(chil)
		}
	}
	dfs(root)

	return res
}

