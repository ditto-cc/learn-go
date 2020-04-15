package bst

import (
	"learn-go/data-structure/utils"
)

func (node *Node) inOrder(visit func(utils.Comparable)) {
	if node == nil {
		return
	}

	node.Lchild.inOrder(visit)
	visit(node.Key)
	node.Rchild.inOrder(visit)
}

func (node *Node) preOrder(visit func(utils.Comparable)) {
	if node == nil {
		return
	}

	visit(node.Key)
	node.Lchild.preOrder(visit)
	node.Rchild.preOrder(visit)
}

func (node *Node) postOrder(visit func(utils.Comparable)) {
	if node == nil {
		return
	}

	node.Lchild.postOrder(visit)
	node.Rchild.postOrder(visit)
	visit(node.Key)
}

// InOrder in-Order traverse
func (bst *BSTree) InOrder(visit func(utils.Comparable)) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}

	bst.root.inOrder(visit)
}

// PreOrder pre-Order traverse
func (bst *BSTree) PreOrder(visit func(utils.Comparable)) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}

	bst.root.preOrder(visit)
}

// PostOrder post-Order traverse
func (bst *BSTree) PostOrder(visit func(utils.Comparable)) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}

	bst.root.postOrder(visit)
}

func (bst *BSTree) InOrderNR(visit func(utils.Comparable)) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}

	s := []*Node{}
	p := bst.root
	for len(s) > 0 || p != nil {
		if p != nil {
			s = append(s, p)
			p = p.Lchild
		} else {
			top := s[len(s)-1]
			s = s[:len(s)-1]
			visit(top.Key)
			p = top.Rchild
		}
	}
}

func (bst *BSTree) PreOrderNR(visit func(utils.Comparable)) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}
	if bst.root == nil {
		return
	}
	s := []*Node{bst.root}
	for len(s) > 0 {
		top := s[len(s)-1]
		s = s[:len(s)-1]
		visit(top.Key)
		if top.Rchild != nil {
			s = append(s, top.Rchild)
		}
		if top.Lchild != nil {
			s = append(s, top.Lchild)
		}
	}
}

func (bst *BSTree) PostOrderNR(visit func(utils.Comparable)) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}
	if bst.root == nil {
		return
	}
	s1 := []*Node{bst.root}
	s2 := []*Node{}
	for len(s1) > 0 {
		top := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		s2 = append(s2, top)
		if top.Lchild != nil {
			s1 = append(s1, top.Lchild)
		}
		if top.Rchild != nil {
			s1 = append(s1, top.Rchild)
		}
	}

	for len(s2) > 0 {
		top := s2[len(s2)-1]
		s2 = s2[:len(s2)-1]
		visit(top.Key)
	}
}

func (bst *BSTree) LevelOrder(visit func(utils.Comparable)) {
	if bst == nil {
		panic("nil BSTree. Error.")
	}
	if bst.root == nil {
		return
	}
	q := []*Node{}
	q = append(q, bst.root)
	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		visit(front.Key)
		if front.Lchild != nil {
			q = append(q, front.Lchild)
		}
		if front.Rchild != nil {
			q = append(q, front.Rchild)
		}
	}
}
