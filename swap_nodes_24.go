package dsa

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	var next, last, temp *ListNode
	for curr != nil && curr.Next != nil {
		next = curr.Next
		temp = next.Next
		// fmt.Println(curr.Val, next.Val, temp.Val)
		next.Next = curr
		curr.Next = temp
		if last == nil {
			head = next
			last = curr
		} else {
			last.Next = next
			last = curr
		}
		// fmt.Println(curr.Val, next.Val, temp.Val, last.Val)
		curr = temp
	}
	return head
}
