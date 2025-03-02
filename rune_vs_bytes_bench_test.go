package dsa

import (
	"testing"
)

func lengthOfLongestSubstringRune(s string) int {
	hash := make(map[rune]int)
	ma, k := 0, 0
	for i, e := range s {
		hash[e] += 1
		for hash[e] > 1 {
			hash[rune(s[k])] -= 1
			k++
		}
		if i-k+1 > ma {
			ma = i - k + 1
		}
	}
	return ma
}

func lengthOfLongestSubstringByte(s string) int {
	maxl := 0
	check := make(map[byte]int)
	j := 0
	for i := range s {
		check[s[i]]++
		for j < i && check[s[i]] > 1 {
			check[s[j]]--
			j++
		}
		currl := i - j + 1
		if currl > maxl {
			maxl = currl
		}
	}
	return maxl
}

func BenchmarkLengthOfLongestSubstringRune(b *testing.B) {
	input := "abcabcbb" // Test input
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstringRune(input)
	}
}

func BenchmarkLengthOfLongestSubstringByte(b *testing.B) {
	input := "abcabcbb" // Test input
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstringByte(input)
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder1(root *TreeNode) [][]int {
	res := make([][]int, 0)
	// check edge case if root is null
	if root == nil {
		return res
	}
	q := make([][]*TreeNode, 0)

	// level 1 traversal
	q = append(q, []*TreeNode{root})
	res = append(res, []int{root.Val})

	for len(q) != 0 {
		level := make([]*TreeNode, 0)
		rlevel := make([]int, 0)
		front := q[0]

		for _, node := range front {
			if node.Left != nil {
				level = append(level, node.Left)
				rlevel = append(rlevel, node.Left.Val)
			}
			if node.Right != nil {
				level = append(level, node.Right)
				rlevel = append(rlevel, node.Right.Val)
			}
		}
		q = q[1:]
		q = append(q, level)
		res = append(res, rlevel)
	}
	return res
}
