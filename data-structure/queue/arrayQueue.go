package queue

import (
	"fmt"
)

// ArrayQueue FIFO
type ArrayQueue struct {
	data []interface{}
}

func CreateArrayQueue() *ArrayQueue {
	return &ArrayQueue{data: []interface{}{}}
}

func (q *ArrayQueue) Empty() bool {
	return q.Size() == 0
}

func (q *ArrayQueue) Size() int {
	return len(q.data)
}

func (q *ArrayQueue) Enqueue(e interface{}) {
	q.data = append(q.data, e)
}

func (q *ArrayQueue) GetFront() interface{} {
	if q.Empty() {
		panic("Empty Queue. Pop Failed.")
	}
	return q.data[0]
}

func (q *ArrayQueue) Dequeue() interface{} {
	if q.Empty() {
		panic("Empty Queue. Pop Failed.")
	}
	ret := q.data[0]
	q.data = q.data[1:]
	return ret
}

func (q *ArrayQueue) String() string {
	str := "Head["
	for i, e := range q.data {
		str += fmt.Sprintf("%v", e)
		if i != q.Size()-1 {
			str += ", "
		}
	}
	return str + "]Tail"
}
