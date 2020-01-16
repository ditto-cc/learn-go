package linkedlist

type DListNode struct {
	Val         int
	Next, Prior *DListNode
}

type DLinkedList struct {
	head *DListNode
	size int
}

func CreateDListNode(val int, prior, next *DListNode) *DListNode {
	return &DListNode{Val: val, Next: next}
}

func CreateDLinkedList() *DLinkedList {
	list := &DLinkedList{size: 0, head: CreateDListNode(0, nil, nil)}
	list.head.Prior = list.head
	list.head.Next = list.head
	return list
}

func (list *DLinkedList) Add(i, e int) {
	p := list.head
	for j := 0; j < i; j++ {
		p = p.Next
	}
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
