package bst

import (
	"learn-go/data-structure/utils"
)

func (node *Node) add(e utils.Comparable) *Node {
	if node == nil {
		return CreateNode(e)
	}

	if r := e.Compare(node.Val); r > 0 {
		node.Rchild = node.Rchild.add(e)
	} else if r < 0 {
		node.Lchild = node.Lchild.add(e)
	}
	return node
}

// Add element to BSTree
func (bst *BSTree) Add(e utils.Comparable) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}
	bst.root = bst.root.add(e)
	bst.size++
}

func (node *Node) getNode(e utils.Comparable) *Node {
	if node == nil {
		return nil
	}

	if r := e.Compare(node.Val); r == 0 {
		return node
	} else if r < 0 {
		return node.Lchild.getNode(e)
	} else {
		return node.Rchild.getNode(e)
	}
}

func (node *Node) removeMin(minNode *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Lchild == nil {
		minNode = node
		return node.Rchild
	}
	node.Lchild = node.Lchild.removeMin(minNode)
	return node
}

func (node *Node) removeMax(maxNode *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Rchild == nil {
		maxNode = node
		return node.Lchild
	}
	node.Rchild = node.Rchild.removeMax(maxNode)
	return node
}

func (bst *BSTree) RemoveMin() utils.Comparable {
	if bst == nil {
		panic("nil BSTree.")
	}

	minNode := bst.root
	bst.root = bst.root.removeMin(minNode)
	bst.size--
	if minNode != nil {
		return minNode.Val
	}
	return nil
}

func (bst *BSTree) RemoveMax() utils.Comparable {
	if bst == nil {
		panic("nil BSTree.")
	}

	maxNode := bst.root
	bst.root = bst.root.removeMax(maxNode)
	bst.size--
	if maxNode != nil {
		return maxNode.Val
	}
	return nil
}

func (node *Node) remove(e utils.Comparable) *Node {
	if node == nil {
		return nil
	}

	retNode := node
	if r := e.Compare(node.Val); r < 0 {
		node.Lchild = node.Lchild.remove(e)
	} else if r > 0 {
		node.Rchild = node.Rchild.remove(e)
	} else {
		if node.Lchild == nil {
			retNode = node.Rchild
			node.Rchild = nil
		} else if node.Rchild == nil {
			retNode = node.Lchild
			node.Lchild = nil
		} else {
			retNode = node.Rchild
			node.Rchild = node.Rchild.removeMin(retNode)
			retNode.Lchild, retNode.Rchild = node.Lchild, node.Rchild
			node.Lchild, node.Rchild = nil, nil
		}
	}
	return retNode
}

func (bst *BSTree) Remove(e utils.Comparable) {
	if bst == nil {
		panic("nil BSTree.")
	}

	bst.root = bst.root.remove(e)
	bst.size--
}

func (bst *BSTree) Contains(e utils.Comparable) bool {
	if bst == nil {
		return false
	}

	return bst.root.getNode(e) != nil
}

func (bst *BSTree) Size() int {
	return bst.size
}

func (bst *BSTree) Empty() bool {
	return bst.size > 0
}

func (node *Node) height() int {
	if node == nil {
		return 0
	}

	return 1 + utils.Max(node.Lchild.height(), node.Rchild.height())
}

func (bst *BSTree) Height() int {
	if bst == nil {
		panic("nil BSTree. Error.")
	}
	return bst.root.height()
}
