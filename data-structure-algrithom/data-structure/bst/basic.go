package bst

func (node *Node) add(e int) *Node {
	if node == nil {
		return newNode(e)
	}
	if e < node.Val {
		node.Lchild = node.Lchild.add(e)
	} else if e > node.Val {
		node.Rchild = node.Rchild.add(e)
	}
	return node
}

// Add element to BSTree
func (bst *BSTree) Add(e int) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}
	bst.root = bst.root.add(e)
	bst.size++
}
