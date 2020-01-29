package bst

func (node *Node) add(e int) *Node {
	if node == nil {
		return CreateNode(e)
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

func (node *Node) getMaxNode() *Node {
	if node == nil {
		return nil
	}
	if node.Rchild == nil {
		return node
	}
	return node.Rchild.getMaxNode()
}

func (node *Node) getMinNode() *Node {
	if node == nil {
		return nil
	}
	if node.Lchild == nil {
		return node
	}
	return node.Lchild.getMinNode()
}

func (node *Node) removeMax() *Node {
	if node == nil {
		return nil
	}
	if node.Lchild == nil {
		return node.Rchild
	}
	node.Lchild = node.Lchild.removeMin()
	return node
}

func (node *Node) removeMin() *Node {
	if node == nil {
		return nil
	}
	if node.Rchild == nil {
		return node.Lchild
	}
	node.Rchild = node.Rchild.removeMin()
	return node
}

func (bst *BSTree) RemoveMin() int {
	var ret int
	if bst == nil && bst.size == 0 {
		panic("Empty BSTree")
	}
	return ret
}

func (bst *BSTree) RemoveMax() int {
	var ret int
	return ret
}
