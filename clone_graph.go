package dsa

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make([]*Node, 100)
	stack := make([]*Node, 0)
	stack = append(stack, node)
	newNode := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	visited[node.Val] = newNode
	for len(stack) != 0 {
		old := stack[0]
		stack = stack[1:]
		new := visited[old.Val]
		for _, oldN := range old.Neighbors {
			var newN *Node
			if visited[oldN.Val] == nil {
				newN = &Node{Val: oldN.Val, Neighbors: make([]*Node, 0)}
				visited[oldN.Val] = newN
				stack = append(stack, oldN)
			} else {
				newN = visited[oldN.Val]
			}
			new.Neighbors = append(new.Neighbors, newN)
		}
	}

	// verify
	stack = append(stack, newNode)
	v2 := make([]bool, 100)
	fmt.Print("[")
	for len(stack) != 0 {
		top := stack[0]
		v2[top.Val] = true
		stack = stack[1:]
		n := make([]int, len(top.Neighbors))
		for i, e := range top.Neighbors {
			n[i] = e.Val
			if !v2[e.Val] {
				stack = append(stack, e)
			}
		}
		fmt.Print(n)
	}
	fmt.Print("]")
	// verify

	return newNode
}
