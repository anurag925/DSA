package dsa

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	res = append(res, []int{root.Val})
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		level := make([]*TreeNode, 0)
		palinLevel := make([]int, 0)
		fmt.Println(queue, res)
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
		queue = level
		res = append(res, palinLevel)
	}

	return res[:len(res)-1]
}
