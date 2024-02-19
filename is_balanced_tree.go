package dsa

import "math"

func checkBalanced(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := checkBalanced(root.Left)
	right := checkBalanced(root.Right)
	if left == -1 || right == -1 || math.Abs(float64(left)-float64(right)) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func isBalanced(root *TreeNode) bool {
	return checkBalanced(root) != -1
}
