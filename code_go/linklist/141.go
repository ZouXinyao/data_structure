package linklist

// 快慢指针
func hasCycle(head *ListNode) bool {
	s, f := head, head
	for s != nil && f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if f == s { return true }
	}
	return false
}
