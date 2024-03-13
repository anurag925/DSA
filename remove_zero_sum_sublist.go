package dsa

import "fmt"

// edge cases ka badsaah

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	if head.Val == 0 {
		head = head.Next
	}

	ahead := head
	nodeHash := make(map[int]*ListNode)
	repeat := true
	last := 0
	for repeat {
		repeat = false
		for ahead != nil {
			last += ahead.Val
			if last == 0 {
				head = ahead.Next
				ahead = head
				repeat = true
				continue
			}
			fmt.Println(last, nodeHash)
			if val, ok := nodeHash[last]; ok {
				fmt.Println(val.Val, ahead.Val)
				for val.Next != ahead {
					val.Next = val.Next.Next
				}
				val.Next = val.Next.Next
				repeat = true
			}
			nodeHash[last] = ahead
			ahead = ahead.Next
		}
		if !repeat && last == 0 {
			return nil
		}
		nodeHash = make(map[int]*ListNode)
		last = 0
	}
	return head
}
