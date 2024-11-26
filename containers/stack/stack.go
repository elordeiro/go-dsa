// Package stack provides a stack data structure implementation.
package stack

import (
	"fmt"
	"iter"
	"strings"
)

// Stack is an implementation of a stack LIFO data structure.
type Stack[V any] []V

// New creates a new stack with the given values if any are provided.
func New[V any](vals ...V) *Stack[V] {
	stk := Stack[V](vals)
	return &stk
}

// Push adds the given values to the stack.
func (s *Stack[V]) Push(val ...V) {
	*s = append(*s, val...)
}

// Pop removes and returns the top value of the stack.
func (s *Stack[V]) Pop() V {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

// IsEmpty returns true if the stack is empty.
func (s Stack[V]) IsEmpty() bool {
	return len(s) == 0
}

// Top returns the top value of the stack without removing it.
func (s Stack[V]) Top() V {
	return s[len(s)-1]
}

// Len returns the number of elements in the stack.
func (s Stack[V]) Len() int {
	return len(s)
}

// All returns a iter.Seq[V] of all the elements in the stack.
func (s Stack[V]) All() iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(s[i]) {
				return
			}
		}
	}
}

// Backwards returns a iter.Seq[V] of all the elements in the stack in reverse order.
func (s Stack[V]) Backwards() iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := 0; i < len(s); i++ {
			if !yield(s[i]) {
				return
			}
		}
	}
}

// Drain returns a iter.Seq[V] of all the elements in the stack by popping them.
// The stack will be empty after this operation.
// It returns a single use iterator.
func (s *Stack[V]) Drain() iter.Seq[V] {
	return func(yield func(V) bool) {
		for !s.IsEmpty() {
			if !yield(s.Pop()) {
				return
			}
		}
	}
}

// String returns a string representation of the stack.
func (s Stack[V]) String() string {
	var sb strings.Builder
	sb.WriteString("$[")
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteString(fmt.Sprint(s[i]))
		if i > 0 {
			sb.WriteByte(' ')
		}
	}
	sb.WriteByte(']')
	return sb.String()
}
