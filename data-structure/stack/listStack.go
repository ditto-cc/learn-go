package stack

import linkedlist "learn-go/data-structure/linkedList"

type ListStack struct {
	data *linkedlist.LinkedList
}

func CreateListStack() *ListStack {
	return &ListStack{data: linkedlist.CreateLinkedList()}
}

func (s *ListStack) Size() int {
	return s.data.Size()
}

func (s *ListStack) Empty() bool {
	return 0 == s.Size()
}

func (s *ListStack) Push(e interface{}) {
	s.data.Prepend(e)
}

func (s *ListStack) Pop() interface{} {
	if s.Empty() {
		panic("Empty Stack. Pop Failed.")
	}
	return s.data.PopLeft()
}

func (s *ListStack) Top() interface{} {
	if s.Empty() {
		panic("Empty Stack. Top Failed.")
	}
	return s.data.Get(0)
}

func (s *ListStack) String() string {
	return "Top" + s.data.String() + "Bottom"
}
