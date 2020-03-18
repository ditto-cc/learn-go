package queue

type Queue interface {
	Size() int
	Empty() bool
	Enqueue(interface{})
	Dequeue() interface{}
	GetFront() interface{}
}
