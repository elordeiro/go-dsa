package stack

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
func (s Stack[V]) IsEmpty() bool {
	return len(s) == 0
}

// Peek returns the top value of the stack without removing it
func (s Stack[V]) Peek() V {
	return s[len(s)-1]
}

// Len returns the number of elements in the stack
func (s Stack[V]) Len() int {
	return len(s)
}

// ----------------------------------------------------------------------------
// Utils
// ----------------------------------------------------------------------------

// All returns a iter.Seq[V] of all the elements in the stack.
func (s *Stack[V]) All() iter.Seq[V] {
	return func(yield func(V) bool) {
		for !s.IsEmpty() {
			if !yield(s.Pop()) {
				return
			}
		}
	}
}

// Enumerate returns an iter.Seq2[int, V] of all the elements in the stack
// and their index. The iteration empties the stack
func (s *Stack[V]) Enumerate(start int) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := start
		for !s.IsEmpty() {
			if !yield(i, s.Pop()) {
				return
			}
			i++
		}
	}
}

// String returns a string representation of the stack
func (s Stack[V]) String() string {
	return fmt.Sprint([]V(s))
}
