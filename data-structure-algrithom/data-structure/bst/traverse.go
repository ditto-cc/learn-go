package bst

func (node *Node) inOrder(visit func(int)) {
	if node == nil {
		return
	}

	node.Lchild.inOrder(visit)
	visit(node.Val)
	node.Rchild.inOrder(visit)
}

func (node *Node) preOrder(visit func(int)) {
	if node == nil {
		return
	}

	visit(node.Val)
	node.Lchild.inOrder(visit)
	node.Rchild.inOrder(visit)
}

func (node *Node) postOrder(visit func(int)) {
	if node == nil {
		return
	}

	node.Lchild.inOrder(visit)
	node.Rchild.inOrder(visit)
	visit(node.Val)
}

// InOrder in-Order traverse
func (bst *BSTree) InOrder(visit func(int)) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}

	bst.root.inOrder(visit)
}

// PreOrder pre-Order traverse
func (bst *BSTree) PreOrder(visit func(int)) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}

	bst.root.preOrder(visit)
}

// PostOrder post-Order traverse
func (bst *BSTree) PostOrder(visit func(int)) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}

	bst.root.postOrder(visit)
}