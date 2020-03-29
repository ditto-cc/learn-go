package avl

import (
	"bytes"
	"fmt"
	"learn-go/data-structure/utils"
)

// Node tree node
type Node struct {
	Key            utils.Comparable
	Val            interface{}
	Height         int
	Lchild, Rchild *Node
}

// AVLTree Binary Search Tree
type AVLTree struct {
	root *Node // Root
	size int   // Size of tree node
}

// CreateNode construct Node
func CreateNode(Val utils.Comparable) *Node {
	return &Node{Key: Val}
}

// CreateAVL construct BST
func CreateAVL() *AVLTree {
	return &AVLTree{}
}

func (avl *AVLTree) String() string {
	if avl == nil || avl.root == nil {
		return ""
	}

	space := ""
	for l := len(fmt.Sprintf("%v-%v", avl.root.Key, avl.root.Val)); l > 0; l-- {
		space += "_"
	}

	h, str, height := 0, bytes.Buffer{}, avl.Height()
	q := []*Node{avl.root}
	for len(q) > 0 {
		levelNum := len(q)
		h++
		for i := 0; i < levelNum; i++ {
			front := q[0]
			q = q[1:]
			if front != nil {
				str.WriteString(fmt.Sprintf("%v-%v", front.Key, front.Val))
			} else {
				str.WriteString(space)
			}
			str.WriteByte(' ')
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
		str.WriteByte('\n')
	}

	return str.String()
}
