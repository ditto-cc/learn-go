package avl

import "learn-go/data-structure/compare"

func (node *Node) inOrder(visit func(compare.Comparable, interface{})) {
	if node == nil {
		return
	}

	node.Lchild.inOrder(visit)
	visit(node.Key, node.Val)
	node.Rchild.inOrder(visit)
}

func (node *Node) preOrder(visit func(compare.Comparable, interface{})) {
	if node == nil {
		return
	}

	visit(node.Key, node.Val)
	node.Lchild.preOrder(visit)
	node.Rchild.preOrder(visit)
}

func (node *Node) postOrder(visit func(compare.Comparable, interface{})) {
	if node == nil {
		return
	}

	node.Lchild.postOrder(visit)
	node.Rchild.postOrder(visit)
	visit(node.Key, node.Val)
}

// InOrder in-Order traverse
func (avl *AVLTree) InOrder(visit func(compare.Comparable, interface{})) {
	if avl == nil {
		panic("nil AVLTree. Error.")
	}

	avl.root.inOrder(visit)
}

// PreOrder pre-Order traverse
func (avl *AVLTree) PreOrder(visit func(compare.Comparable, interface{})) {
	if avl == nil {
		panic("nil AVLTree. Error.")
	}

	avl.root.preOrder(visit)
}

// PostOrder post-Order traverse
func (avl *AVLTree) PostOrder(visit func(compare.Comparable, interface{})) {
	if avl == nil {
		panic("nil AVLTree. Error.")
	}

	avl.root.postOrder(visit)
}

func (avl *AVLTree) InOrderNR(visit func(compare.Comparable, interface{})) {
	if avl == nil {
		panic("nil AVLTree. Error.")
	}

	s := []*Node{}
	p := avl.root
	for len(s) > 0 || p != nil {
		if p != nil {
			s = append(s, p)
			p = p.Lchild
		} else {
			top := s[len(s)-1]
			s = s[:len(s)-1]
			visit(top.Key, top.Val)
			p = top.Rchild
		}
	}
}

func (avl *AVLTree) PreOrderNR(visit func(compare.Comparable, interface{})) {
	if avl == nil {
		panic("nil AVLTree. Error.")
	}
	if avl.root == nil {
		return
	}
	s := []*Node{avl.root}
	for len(s) > 0 {
		top := s[len(s)-1]
		s = s[:len(s)-1]
		visit(top.Key, top.Val)
		if top.Rchild != nil {
			s = append(s, top.Rchild)
		}
		if top.Lchild != nil {
			s = append(s, top.Lchild)
		}
	}
}

func (avl *AVLTree) PostOrderNR(visit func(compare.Comparable, interface{})) {
	if avl == nil {
		panic("nil AVLTree. Error.")
	}
	if avl.root == nil {
		return
	}
	s1 := []*Node{avl.root}
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
		visit(top.Key, top.Val)
	}
}

func (avl *AVLTree) LevelOrder(visit func(compare.Comparable, interface{})) {
	if avl == nil {
		panic("nil AVLTree. Error.")
	}
	if avl.root == nil {
		return
	}
	q := []*Node{}
	q = append(q, avl.root)
	for len(q) > 0 {
		front := q[0]
		q = q[1:]
		visit(front.Key, front.Val)
		if front.Lchild != nil {
			q = append(q, front.Lchild)
		}
		if front.Rchild != nil {
			q = append(q, front.Rchild)
		}
	}
}
