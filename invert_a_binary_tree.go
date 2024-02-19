package dsa

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	root.Left, root.Right = root.Right, root.Left
	return root
}
