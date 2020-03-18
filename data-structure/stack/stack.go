package stack

type Stack interface {
	Size() int
	Empty() bool
	Push(interface{})
	Pop() interface{}
	Top() interface{}
}
