package queue

import linkedlist "learn-go/data-structure/linkedList"

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

func (q *ListQueue) Enqueue(e interface{}) {
	q.data.Append(e)
}

func (q *ListQueue) GetFront() interface{} {
	if q.Empty() {
		panic("Empty Queue. GetFront Failed.")
	}

	return q.data.Get(0)
}

func (q *ListQueue) Dequeue() interface{} {
	if q.Empty() {
		panic("Empty Queue. Pop Failed.")
	}

	return q.data.PopLeft()
}

func (q *ListQueue) String() string {
	return "Head" + q.data.String() + "Tail"
}
