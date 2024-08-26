package deque

import (
	"fmt"
	"iter"
)

// Internal node structure
type node[V any] struct {
	val  V
	prev *node[V]
	next *node[V]
}

// Deque is a double-ended queue implementation
type Deque[V any] struct {
	len   int
	front *node[V]
	back  *node[V]
}

// NewDeque creates a new deque. If vals are provided, they are added to the deque,
// and the deque is initialized.
func NewDeque[V any](vals ...V) *Deque[V] {
	dq := &Deque[V]{}
	for _, val := range vals {
		dq.PushBack(val)
	}
	return dq
}

// PushFront adds an element to the front of the deque
func (dq *Deque[V]) PushFront(val V) {
	node := &node[V]{val: val}
	if dq.front == nil {
		dq.front = node
		dq.back = node
	} else {
		node.next = dq.front
		dq.front.prev = node
		dq.front = node
	}
	dq.len++
}

// PushBack adds an element to the back of the deque
func (dq *Deque[V]) PushBack(val V) {
	node := &node[V]{val: val}
	if dq.back == nil {
		dq.back = node
		dq.front = node
	} else {
		node.prev = dq.back
		dq.back.next = node
		dq.back = node
	}
	dq.len++
}

// PopFront removes and returns the element at the front of the deque.
// Panics if the deque is empty.
func (dq *Deque[V]) PopFront() V {
	if dq.front == nil {
		panic("Deque is empty!")
	}
	val := dq.front.val
	dq.front = dq.front.next
	if dq.front != nil {
		dq.front.prev = nil
	} else {
		dq.back = nil
	}
	dq.len--
	return val
}

// PopBack removes and returns the element at the back of the deque.
// Panics if the deque is empty.
func (dq *Deque[V]) PopBack() V {
	if dq.back == nil {
		panic("Deque is empty!")
	}
	val := dq.back.val
	dq.back = dq.back.prev
	if dq.back != nil {
		dq.back.next = nil
	} else {
		dq.front = nil
	}
	dq.len--
	return val
}

// PeekFront returns the element at the front of the deque without removing it.
func (dq *Deque[V]) PeekFront() V {
	return dq.front.val
}

// PeekBack returns the element at the back of the deque without removing it.
func (dq *Deque[V]) PeekBack() V {
	return dq.back.val
}

// IsEmpty returns true if the deque is empty, false otherwise
func (dq *Deque[V]) IsEmpty() bool {
	return dq.len == 0
}

// Len returns the number of elements in the deque
func (dq *Deque[V]) Len() int {
	return dq.len
}

// ----------------------------------------------------------------------------
// Utils
// ----------------------------------------------------------------------------

// All returns an iter.Seq[V] that yields all elements in the deque
// It returns a single use iterator
func (dq *Deque[V]) All() iter.Seq[V] {
	return func(yield func(V) bool) {
		for !dq.IsEmpty() {
			if !yield(dq.PopFront()) {
				return
			}
		}
	}
}

// Backwards returns an iter.Seq[V] that yields all elements in the deque in reverse order
// as if the deque was a stack.
// It returns a single use iterator
func (dq *Deque[V]) Backwards() iter.Seq[V] {
	return func(yield func(V) bool) {
		for !dq.IsEmpty() {
			if !yield(dq.PopBack()) {
				return
			}
		}
	}
}

// Enumerate returns an iter.Seq2[int, V] that yields all elements in the deque
// along with their index starting from the provided start index.
// It returns a single use iterator
func (dq *Deque[V]) Enumerate(start int) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := start
		for !dq.IsEmpty() {
			if !yield(i, dq.PopFront()) {
				return
			}
			i++
		}
	}
}

// EnumerateBackwards returns an iter.Seq2[int, V] that yields all elements in the deque
// along with their index starting from the provided start index in reverse order
// as if the deque was a stack.
// It returns a single use iterator
func (dq *Deque[V]) EnumerateBackwards(start int) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := start
		for !dq.IsEmpty() {
			if !yield(i, dq.PopBack()) {
				return
			}
			i++
		}
	}
}

// String returns a string representation of the deque
func (dq *Deque[V]) String() string {
	str := "<->["
	for v := range dq.All() {
		str += fmt.Sprintf("%v ", v)
	}
	if len(str) > 4 {
		str = str[:len(str)-1]

	}
	str += "]"
	return str
}
