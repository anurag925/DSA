package dsa

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	curr, next := head, head.Next
	for next != nil {
		head.Next = next.Next
		next.Next = curr
		curr = next
		next = head.Next
	}
	return curr
}
