package containers

import (
	"fmt"
	"iter"
)

type Stack[V any] []V // Stack is a stack of any type

// NewStack creates a new stack with the given values if any
// If no values are given, an empty stack stack is created, however,
// the compiler will require the type to be specified.
func NewStack[V any](vals ...V) Stack[V] {
	return Stack[V](vals)
}

// Push adds the given values to the stack
func (s *Stack[V]) Push(val ...V) {
	*s = append(*s, val...)
}

// Pop removes and returns the top value of the stack
func (s *Stack[V]) Pop() V {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

// IsEmpty returns true if the stack is empty
func (s *Stack[V]) IsEmpty() bool {
	return len(*s) == 0
}

// Peek returns the top value of the stack without removing it
func (s *Stack[V]) Peek() V {
	return (*s)[len(*s)-1]
}

// Len returns the number of elements in the stack
func (s *Stack[V]) Len() int {
	return len(*s)
}

// All iterates over all the elements in the stack
// The iteration is done in LIFO order and empties the stack
func (s *Stack[V]) All() iter.Seq[V] {
	return func(yield func(V) bool) {
		for !s.IsEmpty() {
			if !yield(s.Pop()) {
				return
			}
		}
	}
}

// String returns a string representation of the stack
// The String() method is intended to be used for debugging purposes
// only as it needs to consume the stack while creating a copy of it
func (s Stack[V]) String() string {
	str := "["
	temp := NewStack[V]()
	for !s.IsEmpty() {
		v := s.Pop()
		str += fmt.Sprintf("%v ", v)
		temp.Push(v)
	}
	for v := range temp.All() {
		s.Push(v)
	}
	if len(str) > 1 {
		str = str[:len(str)-1]
	}
	str += "]"
	return str
}
