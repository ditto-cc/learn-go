package queue

import linkedlist "learn-go/data-structure-algrithom/data-structure/linkedList"

type ListQueue struct {
	data *linkedlist.DLinkedList
}

func CreateListQueue() *ListQueue {
	return &ListQueue{data: linkedlist.CreateDLinkedList()}
}

func (q *ListQueue) Empty() bool {
	return q.Size() == 0
}

func (q *ListQueue) Size() int {
	return q.data.Size()
}

func (q *ListQueue) Enqueue(e int) {
	q.data.Append(e)
}

func (q *ListQueue) getFront() int {
	if q.Empty() {
		panic("Empty Queue. GetFront Failed.")
	}

	return q.data.Get(0)
}
func (q *ListQueue) Dequeue() int {
	if q.Empty() {
		panic("Empty Queue. Pop Failed.")
	}

	return q.data.PopLeft()
}

func (q *ListQueue) String() string {
	return "head" + q.data.String() + "tail"
}
