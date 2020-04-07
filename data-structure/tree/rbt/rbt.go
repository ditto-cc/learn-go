package rbt

import "learn-go/data-structure/utils"

const (
	RED   = true
	BLACK = false
)

type Node struct {
	Key            utils.Comparable
	Value          interface{}
	Lchild, Rchild *Node
	Color          bool
}

type RBTree struct {
	root *Node
	size int
}

func newNode(key utils.Comparable, value interface{}) *Node {
	return &Node{
		Key:   key,
		Value: value,
		Color: RED,
	}
}

func NewRBTree() *RBTree {
	return &RBTree{}
}
