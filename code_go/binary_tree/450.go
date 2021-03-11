package main

func deleteNode(root *TreeNode, key int) *TreeNode {

	//总共3种情况
	//1、node只有左树，保留左
	//1、node只有右树，保留右
	//1、node有左树&右树，保留右和左
	if root == nil { return root }
	if root.Val < key{
		root.Right = deleteNode(root.Right,key)
	}else if root.Val > key{
		root.Left = deleteNode(root.Left,key)
	}else{
		// 这里的返回值，如果是根节点就直接返回，非根节点返回到上面
		if root.Right == nil{
			root = root.Left
		}else if root.Left == nil{
			root = root.Right
		}else{
			node := root.Right

			for node.Left!=nil{
				node = node.Left
			}
			node.Left = root.Left
			root = root.Right
		}
	}
	return root
}
