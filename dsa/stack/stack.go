package stack

type Stack[E any] []E

func NewStack[E any](vals ...E) Stack[E] {
	return Stack[E](vals)
}

func (s *Stack[E]) Push(val ...E) {
	*s = append(*s, val...)
}

func (s *Stack[E]) Pop() E {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *Stack[E]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[E]) Peek() E {
	return (*s)[len(*s)-1]
}

func (s *Stack[E]) Len() int {
	return len(*s)
}
