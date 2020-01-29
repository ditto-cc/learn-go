package bst

// Node tree node
type Node struct {
	Val            int
	Lchild, Rchild *Node
}

// BSTree Binary Search Tree
type BSTree struct {
	root *Node // Root
	size int   // Size of tree node
}

// CreateNode construct Node
func CreateNode(Val int) *Node {
	return &Node{Val: Val}
}

// CreateBST construct BST
func CreateBST() *BSTree {
	return &BSTree{}
}
