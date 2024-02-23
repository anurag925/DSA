package dsa

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func lastNode(head *ListNode) (*ListNode, int) {
	l := 1
	for head.Next != nil {
		l++
		head = head.Next
	}
	return head, l
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last, l := lastNode(head)
	rotate := k
	if k > l {
		rotate = k % l
	} else if k == l {
		return head
	}
	revR := l - rotate
	count := 1
	node := head
	for count < revR {
		node = node.Next
		count++
	}

	last.Next = head
	head = node.Next
	node.Next = nil
	return head
}
