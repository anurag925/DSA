package dsa

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	l2h := list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	l2t := list2

	head := list1
	var prev *ListNode = nil
	var next *ListNode = nil

	i := 0
	for i != a {
		prev = head
		head = head.Next
		next = head.Next
		i++
	}

	if prev == nil {
		list1 = l2h
	} else {
		prev.Next = l2h
	}

	head = next
	next = head.Next
	for i != b {
		head = head.Next
		next = head.Next
		i++
	}

	l2t.Next = next

	return list1
}
