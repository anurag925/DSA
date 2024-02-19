package dsa

func hasCycle(head *ListNode) bool {
	has := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := has[head]; ok {
			return true
		}
		has[head] = true
		head = head.Next
	}
	return false
}

func hasCycle2p(head *ListNode) bool {
	if head == nil {
		return false
	}
	next := head.Next
	for head != nil && next != nil {
		if head == next {
			return true
		}
		head = head.Next
		if next.Next == nil {
			return false
		}
		next = next.Next.Next
	}
	return false
}
