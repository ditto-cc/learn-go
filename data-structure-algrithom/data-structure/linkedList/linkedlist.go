package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func CreateListNode(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{size: 0, head: CreateListNode(0, nil)}
}

func (list *LinkedList) Add(i, e int) {
	p := list.head
	for j := 0; j < i; j++ {
		p = p.Next
	}
	p.Next = CreateListNode(e, p.Next)
	list.size++
}

func (list *LinkedList) Append(e int) {
	list.Add(list.size, e)
}

func (list *LinkedList) Prepend(e int) {
	list.Add(0, e)
}
