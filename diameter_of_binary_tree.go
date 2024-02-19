package dsa

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var currMax int

func find(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := find(root.Right)
	left := find(root.Left)
	currMax = max(currMax, right+left)
	return max(right, left) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	currMax = 0
	find(root)
	return currMax
}
