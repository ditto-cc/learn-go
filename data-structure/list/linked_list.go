package linkedlist

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func CreateListNode(val interface{}, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{size: 0, head: CreateListNode(0, nil)}
}

func (list *LinkedList) getPrior(i int) *ListNode {
	p := list.head
	for j := 0; j < i; j++ {
		p = p.Next
	}
	return p
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) Add(i int, e interface{}) {
	if i < 0 || i > list.size {
		panic("Illegal Index. Add Failed.")
	}
	p := list.getPrior(i)
	p.Next = CreateListNode(e, p.Next)
	list.size++
}

func (list *LinkedList) Append(e interface{}) {
	list.Add(list.size, e)
}

func (list *LinkedList) Prepend(e interface{}) {
	list.Add(0, e)
}

func (list *LinkedList) Remove(i int) interface{} {
	if i < 0 || i >= list.size {
		panic("Illegal Index. Remove Failed.")
	}
	p := list.getPrior(i)
	delNode := p.Next
	p.Next = delNode.Next
	ret := delNode.Val
	delNode.Next = nil
	list.size--
	return ret
}

func (list *LinkedList) PopLeft() interface{} {
	return list.Remove(0)
}

func (list *LinkedList) PopRight() interface{} {
	return list.Remove(list.size - 1)
}

func (list *LinkedList) Set(i int, e interface{}) {
	if i < 0 || i >= list.size {
		panic("Illegal Index. Set Failed.")
	}
	p := list.getPrior(i)
	p.Next.Val = e
}

func (list *LinkedList) Contains(e interface{}) bool {
	for p := list.head.Next; p != nil; p = p.Next {
		if p.Val == e {
			return true
		}
	}
	return false
}

func (list *LinkedList) Get(i int) interface{} {
	if i < 0 || i >= list.size {
		panic("Illegal Index. Get Failed.")
	}
	p := list.getPrior(i)
	return p.Val
}

func (list *LinkedList) String() string {
	res := bytes.NewBuffer([]byte{'['})
	for p := list.head.Next; p != nil; p = p.Next {
		res.WriteString(fmt.Sprintf("%v", p.Val))
		if p.Next != nil {
			res.Write([]byte{'-', '>'})
		}
	}
	res.WriteByte(']')
	return res.String()
}
