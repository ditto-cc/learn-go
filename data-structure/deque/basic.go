package deque

import "fmt"

func (q *Deque) PushFront(e interface{}) {
	q.data = append(q.data, e)
	q.size++
}

func (q *Deque) PushBack(e interface{}) {
	q.data = append([]interface{}{e}, q.data...)
	q.size++
}

func (q *Deque) PopLeft() interface{} {
	if q.size == 0 {
		panic("Empty Deque. Error.")
	}
	q.size--
	ret := q.data[0]
	q.data = q.data[1:]
	return ret
}

func (q *Deque) PopRight() interface{} {
	if q.size == 0 {
		panic("Empty Deque. Error.")
	}
	q.size--
	ret := q.data[q.size]
	q.data = q.data[:q.size]
	return ret
}

func (q *Deque) Erase(begin, end int) {
	if q.size == 0 {
		panic("Empty Deque. Error.")
	}
	if begin > end {
		panic("begin must lower than end. Error.")
	}
	if begin < 0 || begin >= q.size || end < 0 || end >= q.size {
		panic(fmt.Sprintf("begin, end must between [0, %d)", q.size))
	}
	q.data = append(q.data[0:begin], q.data[end+1:])
	q.size -= end - begin + 1
}

func (q *Deque) Remove(i int) interface{} {
	if q.size == 0 {
		panic("Empty Deque. Error.")
	}
	ret := q.data[i]
	q.data = append(q.data[:i], q.data[i+1:]...)
	q.size--
	return ret
}

func (q *Deque) Get(i int) interface{} {
	if q.size == 0 {
		panic("Empty Deque. Error.")
	}
	return q.data[i]
}
