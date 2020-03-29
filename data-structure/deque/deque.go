package deque

import (
	"bytes"
	"fmt"
)

type Deque struct {
	data []interface{}
	size int
}

func NewDeque(initCap int) *Deque {
	return &Deque{
		data: make([]interface{}, 0, initCap),
		size: 0,
	}
}

func (q *Deque) String() string {
	res := bytes.NewBuffer([]byte{'['})
	for i, e := range q.data {
		res.WriteString(fmt.Sprintf("%v", e))
		if i != q.size-1 {
			res.Write([]byte{',', ' '})
		}
	}
	res.WriteByte(']')
	return res.String()
}
