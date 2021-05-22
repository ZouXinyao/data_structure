package linklist

func swapPairs(head *ListNode) *ListNode {
	res := &ListNode{0, head}
	prev, curr := res, head
	for curr != nil && curr.Next != nil {
		next := curr.Next
		// 互换位置过程
		prev.Next = next
		curr.Next = next.Next
		next.Next = curr
		// 换位置过程
		prev, curr = curr, curr.Next
	}
	return res.Next
}
