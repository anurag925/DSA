package main

type Node struct {
	Right *Node
	Left  *Node
	Val   int
}

type BinaryTree struct {
	Root *Node
}

func main() {
	tree := BinaryTree{}
	tree.Root = &Node{
		Right: nil,
		Left:  nil,
		Val:   5,
	}
}
