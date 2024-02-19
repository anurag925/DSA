package dsa

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			tail.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		tail = tail.Next
	}

	for list1 != nil {
		tail.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		tail = tail.Next
	}

	for list2 != nil {
		tail.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		tail = tail.Next
	}

	return head
}
