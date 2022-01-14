package stack

type Stack struct {
	size    int
	len     int
	head    *Node
	current *Node
}

type Node struct {
	value interface{}
	next  *Node
	prev  *Node
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

	node := &Node{value: v}

	if s.len == 0 {
		s.head = node
		s.current = node
		s.len++
		return
	}

	current := s.head

	for current.next != nil {
		current = current.next
	}

	current.next = node
	node.prev = current
	s.current = node
	s.len++
}

func (s *Stack) Pop() interface{} {
	if s.len == 0 {
		return nil
	}

	current := s.current
	if s.current.prev != nil {
		s.current = s.current.prev
	}
	s.len--

	return current.value
}

func (s *Stack) Size() int {
	return s.len
}

func (s *Stack) Clear() {
	s.len = 0
	s.head = nil
	s.current = nil
}

func (s *Stack) Peek() interface{} {
	if s.len == 0 {
		return nil
	}

	return s.current.value
}
