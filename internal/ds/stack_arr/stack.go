package stack_arr

type Stack struct {
	size int
	len  int
	data []interface{}
}

func NewStack(size int) Stack {
	return Stack{
		size: size,
	}
}

func (s *Stack) Push(v interface{}) {
	if s.len == s.size {
		return
	}

	s.data = append(s.data, v)
	s.len++
}

func (s *Stack) Pop() interface{} {
	if s.len == 0 {
		return nil
	}

	s.len--
	current := s.data[s.len]

	return current
}

func (s *Stack) Size() int {
	return s.len
}

func (s *Stack) Clear() {
	s.len = 0
	s.data = []interface{}{}
}

func (s *Stack) Peek() interface{} {
	if s.len == 0 {
		return nil
	}

	return s.data[s.len]
}
