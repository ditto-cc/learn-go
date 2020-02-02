package stack

import linkedlist "learn-go/data-structure-algrithom/data-structure/linkedList"

type ListStack struct {
	data linkedlist.LinkedList
}

func (s *ListStack) Size() int {
	return s.data.Size()
}

func (s *ListStack) Empty() bool {
	return 0 == s.Size()
}

func (s *ListStack) Push(e int) {
	s.data.Prepend(e)
}

func (s *ListStack) Pop() int {
	if s.Empty() {
		panic("Empty Stack. Pop Failed.")
	}
	return s.data.PopLeft()
}

func (s *ListStack) Top() int {
	if s.Empty() {
		panic("Empty Stack. Top Failed.")
	}
	return s.data.Get(0)
}
