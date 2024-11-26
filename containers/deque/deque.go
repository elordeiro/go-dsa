// Package deque provides a double-ended queue implementation
package deque

import (
	"fmt"
	"iter"
	"strings"
)

// Internal node structure
type node[V any] struct {
	val  V
	prev *node[V]
	next *node[V]
}

// Deque is a double-ended queue implemented as a doubly linked list.
type Deque[V any] struct {
	len   int
	front *node[V]
	back  *node[V]
}

// New creates a new deque. If vals are provided, they are added to the deque,
// and the deque is initialized.
func New[V any](vals ...V) *Deque[V] {
	dq := &Deque[V]{}
	for _, val := range vals {
		dq.PushBack(val)
	}
	return dq
}

// PushFront adds an element to the front of the deque.
func (dq *Deque[V]) PushFront(val ...V) {
	for _, v := range val {
		node := &node[V]{val: v}
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
}

// PushBack adds an element to the back of the deque.
func (dq *Deque[V]) PushBack(val ...V) {
	for _, v := range val {
		node := &node[V]{val: v}
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
}

// PopFront removes and returns the element at the front of the deque.
// Panics if the deque is empty.
func (dq *Deque[V]) PopFront() V {
	if dq.front == nil {
		panic("attempt to pop from an empty deque.\n\tfunc: deque.PopFront()")
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
		panic("attempt to pop from an empty deque.\n\tfunc: deque.PopBack()")
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

// Front returns the element at the front of the deque without removing it.
// Panics if the deque is empty.
func (dq Deque[V]) Front() V {
	if dq.front == nil {
		panic("attempt to access an empty deque.\n\tfunc: deque.Front()")
	}
	return dq.front.val
}

// Back returns the element at the back of the deque without removing it.
// Panics if the deque is empty.
func (dq Deque[V]) Back() V {
	if dq.back == nil {
		panic("attempt to access an empty deque.\n\tfunc: deque.Back()")
	}
	return dq.back.val
}

// IsEmpty returns true if the deque is empty, false otherwise.
func (dq Deque[V]) IsEmpty() bool {
	return dq.len == 0
}

// Len returns the number of elements in the deque.
func (dq Deque[V]) Len() int {
	return dq.len
}

// All returns an iter.Seq[V] that yields all elements in the deque in FIFO order.
func (dq Deque[V]) All() iter.Seq[V] {
	return func(yield func(V) bool) {
		for cur := dq.front; cur != nil; cur = cur.next {
			if !yield(cur.val) {
				return
			}
		}
	}
}

// Backwards returns an iter.Seq[V] that yields all elements in the deque in LIFO order.
func (dq Deque[V]) Backwards() iter.Seq[V] {
	return func(yield func(V) bool) {
		for cur := dq.back; cur != nil; cur = cur.prev {
			if !yield(cur.val) {
				return
			}
		}
	}
}

// Drain returns an iter.Seq[V] that yields all elements in the deque in FIFO
// order by popping elements from the front of the deque. The deque is emptied
// after calling Drain.
// It returns a single use iterator.
func (dq *Deque[V]) Drain() iter.Seq[V] {
	return func(yield func(V) bool) {
		for !dq.IsEmpty() {
			if !yield(dq.PopFront()) {
				return
			}
		}
	}
}

// DrainBackwards returns an iter.Seq[V] that yields all elements in the deque
// in LIFO order by popping elements from the back of the deque. The deque is
// emptied after calling DrainBackwards.
// It returns a single use iterator.
func (dq *Deque[V]) DrainBackwards() iter.Seq[V] {
	return func(yield func(V) bool) {
		for !dq.IsEmpty() {
			if !yield(dq.PopBack()) {
				return
			}
		}
	}
}

// String returns a string representation of the deque.
func (dq Deque[V]) String() string {
	var sb strings.Builder
	sb.WriteByte('<')
	cur := dq.front
	for cur != nil {
		sb.WriteString(fmt.Sprint(cur.val))
		if cur.next != nil {
			sb.WriteByte(' ')
		}
		cur = cur.next
	}
	sb.WriteByte('>')
	return sb.String()
}
