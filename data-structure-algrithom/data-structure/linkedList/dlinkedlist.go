package linkedlist

import "strconv"

type DListNode struct {
	Val         int
	Next, Prior *DListNode
}

type DLinkedList struct {
	head *DListNode
	size int
}

func CreateDListNode(val int, prior, next *DListNode) *DListNode {
	return &DListNode{Val: val, Prior: prior, Next: next}
}

func CreateDLinkedList() *DLinkedList {
	list := &DLinkedList{size: 0, head: CreateDListNode(0, nil, nil)}
	list.head.Prior = list.head
	list.head.Next = list.head
	return list
}

func (list *DLinkedList) getPrior(i int) *DListNode {
	p := list.head
	for j := 0; j < i; j++ {
		p = p.Next
	}
	return p
}

func (list *DLinkedList) Add(i, e int) {
	if i < 0 || i > list.size {
		panic("Illegal Index. Add Failed.")
	}
	p := list.getPrior(i)
	newNode := CreateDListNode(e, p, p.Next)
	newNode.Next.Prior = newNode
	newNode.Prior.Next = newNode
	list.size++
}

func (list *DLinkedList) Append(e int) {
	list.Add(list.size, e)
}

func (list *DLinkedList) Prepend(e int) {
	list.Add(0, e)
}

func (list *DLinkedList) Remove(i int) int {
	if i < 0 || i >= list.size {
		panic("Illegal Index. Remove Failed.")
	}
	p := list.getPrior(i)
	delNode := p.Next
	delNode.Prior.Next = delNode.Next
	delNode.Next.Prior = delNode.Prior
	delNode.Next = nil
	delNode.Prior = nil
	list.size--
	return delNode.Val
}

func (list *DLinkedList) PopLeft() int {
	return list.Remove(0)
}

func (list *DLinkedList) PopRight() int {
	return list.Remove(list.size - 1)
}

func (list *DLinkedList) Set(i, e int) {
	p := list.getPrior(i)
	p.Next.Val = e
}

func (list *DLinkedList) Get(i int) int {
	if i < 0 || i >= list.size {
		panic("Illegal Index. Get Failed.")
	}
	p := list.getPrior(i)
	return p.Next.Val
}

func (list *DLinkedList) Contains(e int) bool {
	for p := list.head.Next; p != nil; p = p.Next {
		if p.Val == e {
			return true
		}
	}
	return false
}

func (list *DLinkedList) Size() int {
	return list.size
}

func (list *DLinkedList) String() string {
	str := "[<->"
	for p := list.head.Next; p != list.head; p = p.Next {
		str += strconv.Itoa(p.Val)
		if list.head != p.Next {
			str += "<->"
		}
	}
	return str + "<->]"
}
