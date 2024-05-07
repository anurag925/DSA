package dsa

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	s := make([]*ListNode, 0)
	for head != nil {
		for len(s) > 0 && s[len(s)-1].Val < head.Val {
			s = s[:len(s)-1]
		}
		s = append(s, head)
		head = head.Next
	}
	newHead := s[0]
	curr := newHead
	for _, node := range s[1:] {
		curr.Next = node
		curr = node
	}
	return newHead
}
