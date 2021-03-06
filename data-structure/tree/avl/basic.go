package avl

import (
	"learn-go/data-structure/utils"
)

func (node *Node) height() int {
	if node == nil {
		return 0
	}
	return node.Height
}

func (avl *Tree) Height() int {
	if avl == nil {
		panic("nil AVLTree.")
	}
	return avl.root.height()
}

func (node *Node) getBalanceFactor() int {
	if node == nil {
		return 0
	}
	return node.Lchild.height() - node.Rchild.height()
}

func (node *Node) leftRotate() *Node {
	if node == nil {
		return nil
	}

	rc := node.Rchild
	node.Rchild = rc.Lchild
	rc.Lchild = node

	node.Height = 1 + utils.MaxInt(node.Lchild.height(), node.Rchild.height())
	rc.Height = 1 + utils.MaxInt(rc.Lchild.height(), rc.Rchild.height())
	return rc
}

func (node *Node) rightRotate() *Node {
	if node == nil {
		return nil
	}

	lc := node.Lchild
	node.Lchild = lc.Rchild
	lc.Rchild = node

	node.Height = 1 + utils.MaxInt(node.Lchild.height(), node.Rchild.height())
	lc.Height = 1 + utils.MaxInt(lc.Lchild.height(), lc.Rchild.height())
	return lc
}

func (node *Node) rotate() *Node {
	if node == nil {
		return nil
	}

	node.Height = 1 + utils.MaxInt(node.Lchild.height(), node.Rchild.height())
	balanceFactor := node.getBalanceFactor()

	if balanceFactor < 2 && balanceFactor > -2 {
		return node
	}

	lb, rb := node.Lchild.getBalanceFactor(), node.Rchild.getBalanceFactor()
	switch {
	case 1 < balanceFactor && 0 <= lb:
		return node.rightRotate()
	case -1 > balanceFactor && 0 >= rb:
		return node.leftRotate()
	case 1 < balanceFactor && 0 > lb:
		node.Lchild = node.Lchild.leftRotate()
		return node.rightRotate()
	case -1 > balanceFactor && 0 < rb:
		node.Rchild = node.Rchild.rightRotate()
		return node.leftRotate()
	}
	return node
}

func (avl *Tree) add(node *Node, key utils.Comparable, val interface{}) *Node {
	if node == nil {
		avl.size++
		return NewNode(key, val)
	}

	if r := key.Compare(node.Key); r > 0 {
		node.Rchild = avl.add(node.Rchild, key, val)
	} else if r < 0 {
		node.Lchild = avl.add(node.Lchild, key, val)
	} else {
		node.Val = val
	}

	return node.rotate()
}

func (avl *Tree) Add(key utils.Comparable, val interface{}) {
	avl.root = avl.add(avl.root, key, val)
}

func (node *Node) getNode(key utils.Comparable) *Node {
	if node == nil {
		return nil
	}

	if r := key.Compare(node.Key); r > 0 {
		return node.Rchild.getNode(key)
	} else if r < 0 {
		return node.Lchild.getNode(key)
	} else {
		return node
	}
}

func (node *Node) removeMin(minNode *Node) *Node {
	if node == nil {
		return nil
	}

	retNode := node
	if node.Lchild == nil {
		retNode = node.Rchild
		node.Rchild = nil
		minNode = node
	} else {
		retNode.Lchild = retNode.Lchild.removeMin(minNode)
	}
	return retNode.rotate()
}

func (node *Node) removeMax(maxNode *Node) *Node {
	if node == nil {
		return nil
	}

	retNode := node
	if node.Rchild == nil {
		retNode = node.Rchild
		node.Lchild = nil
		maxNode = node
	} else {
		retNode.Rchild = retNode.Rchild.removeMax(maxNode)
	}
	return retNode.rotate()
}

func (avl *Tree) RemoveMin() interface{} {
	if avl == nil {
		return nil
	}
	minNode := avl.root
	avl.root = avl.root.removeMin(minNode)
	if minNode == nil {
		return nil
	}
	avl.size--
	return minNode.Val
}

func (avl *Tree) RemoveMax() interface{} {
	if avl == nil {
		return nil
	}
	maxNode := avl.root
	avl.root = avl.root.removeMax(maxNode)
	if maxNode == nil {
		return nil
	}
	avl.size--
	return maxNode.Val
}

func (avl *Tree) remove(node *Node, key utils.Comparable) *Node {
	if node == nil {
		return nil
	}

	retNode := node
	if r := key.Compare(node.Key); r > 0 {
		retNode.Rchild = avl.remove(retNode.Rchild, key)
	} else if r < 0 {
		retNode.Lchild = avl.remove(retNode.Lchild, key)
	} else {
		if node.Lchild == nil {
			retNode = node.Rchild
			node.Rchild = nil
		} else if node.Rchild == nil {
			retNode = node.Lchild
			node.Lchild = nil
		} else {
			var successor *Node = node.Rchild
			node.Rchild = node.removeMin(successor)
			successor.Lchild, successor.Rchild = node.Lchild, node.Rchild
			node.Lchild, node.Rchild = nil, nil
			retNode = successor
		}
	}
	return retNode.rotate()
}

func (avl *Tree) Remove(key utils.Comparable) interface{} {
	if avl == nil {
		panic("nil AVLTree.")
	}
	if node := avl.root.getNode(key); node == nil {
		return nil
	} else {
		avl.root = avl.remove(avl.root, key)
		avl.size--
		return node.Val
	}
}

func (avl *Tree) Contains(key utils.Comparable) bool {
	if avl == nil {
		panic("nil AVLTree.")
	}
	return avl.root.getNode(key) != nil
}

func (avl *Tree) Get(key utils.Comparable) interface{} {
	if avl == nil {
		panic("nil AVLTree.")
	}

	if node := avl.root.getNode(key); node != nil {
		return node.Val
	}
	return nil
}

func (avl *Tree) Set(key utils.Comparable, val interface{}) bool {
	if avl == nil {
		panic("nil Tree")
	}

	if node := avl.root.getNode(key); node != nil {
		node.Val = val
		return true
	}
	return false
}

func (avl *Tree) Size() int {
	return avl.size
}

func (avl *Tree) Empty() bool {
	return avl.size > 0
}
