package stack

type ArrayStack struct {
	data []int
}

func CreateArrayStack() *ArrayStack {
	return &ArrayStack{data: []int{}}
}

func (s *ArrayStack) Size() int {
	return len(s.data)
}

func (s *ArrayStack) Empty() bool {
	return 0 == s.Size()
}

func (s *ArrayStack) Push(e int) {
	s.data = append(s.data, e)
}

func (s *ArrayStack) Pop() int {
	if s.Empty() {
		panic("Empty Stack. Pop Failed.")
	}
	tail := s.Size() - 1
	ret := s.data[tail]
	s.data = s.data[:tail]
	return ret
}

func (s *ArrayStack) Top() int {
	if s.Empty() {
		panic("Empty Stack. Get Failed.")
	}
	return s.data[s.Size()-1]
}
