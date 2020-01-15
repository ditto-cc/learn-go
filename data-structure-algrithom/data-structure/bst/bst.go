package bst

// Node tree node
type Node struct {
	Val            int
	Lchild, Rchild *Node
}

// BSTree Binary Search Tree
type BSTree struct {
	root *Node // Root
	size int   // Size of treenode
}

// NewNode construct Node
func newNode(Val int) *Node {
	return &Node{Val: Val}
}

// NewBSTree construct BST
func NewBSTree() *BSTree {
	return &BSTree{}
}
