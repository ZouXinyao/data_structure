package linklist

func detectCycle(head *ListNode) *ListNode {
	s, f := head, head
	for s != nil && f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		// 当判断为环时，需要做的就是画个图，s距交点的举例=head到交点的举例
		if f == s {
			f = head
			for f != s {
				s, f = s.Next, f.Next
			}
			return f
		}
	}
	return nil
}
