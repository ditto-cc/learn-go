package avl

import (
	"fmt"
)

type Comparable interface {
	Compare(Comparable) int
}

// Node tree node
type Node struct {
	Val            Comparable
	BalanceFactor  int
	Lchild, Rchild *Node
}

// AVLTree Binary Search Tree
type AVLTree struct {
	root *Node // Root
	size int   // Size of tree node
}

// CreateNode construct Node
func CreateNode(Val Comparable) *Node {
	return &Node{Val: Val}
}

// CreateBST construct BST
func CreateBST() *AVLTree {
	return &AVLTree{}
}

func (avl *AVLTree) String() string {
	if avl == nil || avl.root == nil {
		return ""
	}
	var str string
	height := avl.Height()
	h := 0
	q := []*Node{avl.root}
	for len(q) > 0 {
		levelNum := len(q)
		h++
		for i := 0; i < levelNum; i++ {
			front := q[0]
			q = q[1:]
			if front != nil {
				str += fmt.Sprintf("%s", front.Val) + " "
			} else {
				str += "  "
			}
			if h != height {
				if front == nil {
					q = append(q, nil)
					q = append(q, nil)
				} else {
					q = append(q, front.Lchild)
					q = append(q, front.Rchild)
				}
			}
		}
		str += "\n"
	}

	return str
}
