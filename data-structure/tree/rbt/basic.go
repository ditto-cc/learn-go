package rbt

import "learn-go/data-structure/utils"

func (node *Node) isRed() bool {
	if node == nil {
		return false
	}
	return node.Color == RED
}

func (node *Node) getNode(key utils.Comparable) *Node {
	if node == nil {
		return nil
	}
	r := key.Compare(node.Key)
	if r < 0 {
		return node.Lchild.getNode(key)
	} else if r > 0 {
		return node.Rchild.getNode(key)
	}
	return node
}

func (tree *RBTree) Get(key utils.Comparable) interface{} {
	node := tree.root.getNode(key)
	if node == nil {
		return nil
	}
	return node.Value
}

func (tree *RBTree) Contains(key utils.Comparable) bool {
	return tree.root.getNode(key) != nil
}

func (node *Node) leftRotate() *Node {
	rc := node.Rchild

	node.Rchild = rc.Lchild
	rc.Lchild = node

	rc.Color = node.Color
	node.Color = RED
	return rc
}

func (node *Node) flipColor() {
	node.Color = RED
	node.Lchild.Color = BLACK
	node.Rchild.Color = BLACK
}

func (node *Node) rightRotate() *Node {
	lc := node.Lchild

	node.Lchild = lc.Rchild
	lc.Rchild = node

	lc.Color = node.Color
	node.Color = RED

	return node
}

func (tree *RBTree) add(node *Node, key utils.Comparable, value interface{}) *Node {
	if node == nil {
		tree.size++
		return newNode(key, value)
	}

	if r := key.Compare(node.Key); r < 0 {
		node.Lchild = tree.add(node.Lchild, key, value)
	} else if r > 0 {
		node.Rchild = tree.add(node.Rchild, key, value)
	} else {
		node.Value = value
	}

	if node.Rchild.isRed() && !node.Lchild.isRed() {
		node = node.leftRotate()
	}

	if node.Lchild.isRed() && node.Lchild.Lchild.isRed() {
		node = node.rightRotate()
	}

	if node.Lchild.isRed() && node.Rchild.isRed() {
		node.flipColor()
	}

	return node
}

func (tree *RBTree) Add(key utils.Comparable, value interface{}) {
	tree.root = tree.add(tree.root, key, value)
	tree.root.Color = BLACK
}
