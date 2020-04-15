package bst

import (
	"learn-go/data-structure/utils"
)

func (bst *BSTree) add(node *Node, key utils.Comparable, val interface{}) *Node {
	if node == nil {
		bst.size++
		return NewNode(key, val)
	}

	if r := key.Compare(node.Key); r > 0 {
		node.Rchild = bst.add(node.Rchild, key, val)
	} else if r < 0 {
		node.Lchild = bst.add(node.Lchild, key, val)
	} else {
		node.Val = val
	}
	return node
}

// Add element to BSTree
func (bst *BSTree) Add(key utils.Comparable, val interface{}) {
	bst.root = bst.add(bst.root, key, val)
}

func (node *Node) getNode(key utils.Comparable) *Node {
	if node == nil {
		return nil
	}

	if r := key.Compare(node.Key); r == 0 {
		return node
	} else if r < 0 {
		return node.Lchild.getNode(key)
	} else {
		return node.Rchild.getNode(key)
	}
}

func (bst *BSTree) removeMin(node, minNode *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Lchild == nil {
		minNode = node
		bst.size--
		return node.Rchild
	}
	node.Lchild = bst.removeMin(node.Lchild, minNode)
	return node
}

func (bst *BSTree) removeMax(node *Node, maxNode *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Rchild == nil {
		maxNode = node
		bst.size--
		return node.Lchild
	}
	node.Rchild = bst.removeMax(node.Rchild, maxNode)
	return node
}

func (bst *BSTree) RemoveMin() utils.Comparable {
	if bst == nil {
		panic("nil BSTree.")
	}

	var minNode *Node
	bst.root = bst.removeMin(bst.root, minNode)
	if minNode != nil {
		return minNode.Key
	}
	return nil
}

func (bst *BSTree) RemoveMax() utils.Comparable {
	if bst == nil {
		panic("nil BSTree.")
	}

	var maxNode *Node
	bst.root = bst.removeMax(bst.root, maxNode)
	if maxNode != nil {
		return maxNode.Key
	}
	return nil
}

func (bst *BSTree) remove(node *Node, key utils.Comparable) *Node {
	if node == nil {
		return nil
	}

	retNode := node
	if r := key.Compare(node.Key); r < 0 {
		node.Lchild = bst.remove(node.Lchild, key)
	} else if r > 0 {
		node.Rchild = bst.remove(node.Rchild, key)
	} else {
		if node.Lchild == nil {
			retNode = node.Rchild
			node.Rchild = nil
		} else if node.Rchild == nil {
			retNode = node.Lchild
			node.Lchild = nil
		} else {
			retNode = node.Rchild
			node.Rchild = bst.removeMin(bst.root, retNode)
			retNode.Lchild, retNode.Rchild = node.Lchild, node.Rchild
			node.Lchild, node.Rchild = nil, nil
		}
	}
	return retNode
}

func (bst *BSTree) Remove(key utils.Comparable) {
	if bst == nil {
		panic("nil BSTree.")
	}

	bst.root = bst.remove(bst.root, key)
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

	return 1 + utils.MaxInt(node.Lchild.height(), node.Rchild.height())
}

func (bst *BSTree) Height() int {
	if bst == nil {
		panic("nil BSTree. Error.")
	}
	return bst.root.height()
}
