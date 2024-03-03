package dsa

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func checkPalin(level []int) bool {
	start, end := 0, len(level)-1
	for start <= end {
		if level[start] != level[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var level []*TreeNode
	var palinLevel []int
	for len(queue) != 0 {
		level = make([]*TreeNode, 0)
		palinLevel = make([]int, 0)

		for _, e := range queue {
			if e.Left != nil {
				level = append(level, e.Left)
				palinLevel = append(palinLevel, e.Left.Val)
			}
			if e.Right != nil {
				level = append(level, e.Right)
				palinLevel = append(palinLevel, e.Right.Val)
			}
		}
		if !checkPalin(palinLevel) {
			return false
		}
		queue = level
	}

	return true
}
