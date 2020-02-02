package queue

import "strconv"

// ArrayQueue FIFO
type ArrayQueue struct {
	data []int
}

func CreateArrayQueue() *ArrayQueue {
	return &ArrayQueue{data: []int{}}
}

func (q *ArrayQueue) Empty() bool {
	return q.Size() == 0
}

func (q *ArrayQueue) Size() int {
	return len(q.data)
}

func (q *ArrayQueue) Enqueue(e int) {
	q.data = append(q.data, e)
}

func (q *ArrayQueue) GetFront() int {
	if q.Empty() {
		panic("Empty Queue. Pop Failed.")
	}
	return q.data[0]
}

func (q *ArrayQueue) Dequeue() int {
	if q.Empty() {
		panic("Empty Queue. Pop Failed.")
	}
	ret := q.data[0]
	q.data = q.data[1:]
	return ret
}

func (q *ArrayQueue) String() string {
	str := "head["
	for i, e := range q.data {
		str += strconv.Itoa(e)
		if i != q.Size()-1 {
			str += ", "
		}
	}
	return str + "]tail"
}
