package stack

type stack struct {
	stack []any
	len   int
	cap   int
}

func Stack() *stack {
	return &stack{
		stack: make([]any, 4, 4),
		len:   0,
		cap:   4,
	}
}

func (s *stack) Push(val any) {
	if s.len == s.cap {
		s.stack = append(s.stack, make([]any, s.len, s.cap)...)
		s.cap *= 2
	}
	s.stack[s.len] = val
	s.len++
}

func (s *stack) Pop() any {
	res := s.stack[s.len-1]
	s.len--
	if s.len < s.cap/4 {
		s.stack = s.stack[:s.cap/2]
		s.cap /= 2
	}
	return res
}

func (s *stack) IsEmpty() bool {
	return s.len == 0
}

func (s *stack) Peek() any {
	return s.stack[s.len-1]
}
