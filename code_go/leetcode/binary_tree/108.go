package main

// 平衡二叉树，直接从数组中间开始分别向左，向右遍历
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums) - 1)
}

func helper(nums []int, l, r int) *TreeNode {
	if l > r { return nil }

	m := l + (r - l) / 2
	node := &TreeNode{Val: nums[m]}
	node.Left = helper(nums, l, m - 1)
	node.Right = helper(nums, m + 1, r)
	return node
}
