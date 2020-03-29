package bst

import (
	"fmt"
	"learn-go/data-structure/utils"
)

// Node tree node
type Node struct {
	Val            utils.Comparable
	Lchild, Rchild *Node
}

// BSTree Binary Search Tree
type BSTree struct {
	root *Node // Root
	size int   // Size of tree node
}

// CreateNode construct Node
func CreateNode(Val utils.Comparable) *Node {
	return &Node{Val: Val}
}

// CreateBST construct BST
func CreateBST() *BSTree {
	return &BSTree{}
}

func (bst *BSTree) String() string {
	if bst == nil || bst.root == nil {
		return ""
	}
	var str string
	height := bst.Height()
	h := 0
	q := []*Node{bst.root}
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
