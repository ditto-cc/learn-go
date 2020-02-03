package stack

import "fmt"

type ArrayStack struct {
	data []interface{}
}

func CreateArrayStack() *ArrayStack {
	return &ArrayStack{data: []interface{}{}}
}

func (s *ArrayStack) Size() int {
	return len(s.data)
}

func (s *ArrayStack) Empty() bool {
	return 0 == s.Size()
}

func (s *ArrayStack) Push(e interface{}) {
	s.data = append(s.data, e)
}

func (s *ArrayStack) Pop() interface{} {
	if s.Empty() {
		panic("Empty Stack. Pop Failed.")
	}
	tail := s.Size() - 1
	ret := s.data[tail]
	s.data = s.data[:tail]
	return ret
}

func (s *ArrayStack) Top() interface{} {
	if s.Empty() {
		panic("Empty Stack. Get Failed.")
	}
	return s.data[s.Size()-1]
}

func (s *ArrayStack) String() string {
	str := "Bottom["
	for i := 0; i < s.Size(); i++ {
		str += fmt.Sprintf("%v", s.data[i])
		if i != s.Size()-1 {
			str += ", "
		}
	}
	return str + "]Top"
}
