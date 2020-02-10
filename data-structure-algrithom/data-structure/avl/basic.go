package avl

import "math"

func (node *Node) height() int {
	if node == nil {
		return 0
	}
	return int(math.Max(float64(node.Lchild.height()), float64(node.Rchild.height()))) + 1
}

func (avl *AVLTree) Height() int {
	if avl == nil || avl.root == nil {
		return 0
	}
	return avl.root.height()
}