package list

import (
	"fmt"
	"iter"
)

// List is a singly linked list implementation
type List[V any] struct {
	Val  V
	Next *List[V]
}

// ----------------------------------------------------------------------------

// NewList creates a new list. If vals are provided, they are added to the list,
// and the list is initialized.
func NewList[V any](vals ...V) *List[V] {
	head := &List[V]{}
	curr := head
	for _, val := range vals {
		curr.Next = &List[V]{val, nil}
		curr = curr.Next
	}
	return head.Next
}

// Append adds a new value to the end of the list and returns the new list
func (l *List[V]) Append(val V) *List[V] {
	if l == nil {
		return &List[V]{val, nil}
	}
	head := l
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &List[V]{val, nil}
	return head
}

// Prepend adds a new value to the beginning of the list and returns the new list
func (l *List[V]) Prepend(val V) *List[V] {
	return &List[V]{val, l}
}

// Remove removes the first occurrence of the value from the list and returns the new list
func Remove[V comparable](list *List[V], val V) *List[V] {
	head := &List[V]{}
	head.Next = list
	prev := head
	curr := list
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
			break
		}
		prev = curr
		curr = curr.Next
	}
	return head.Next
}

// RemoveAll removes all occurrences of the value from the list and returns the new list
func RemoveAll[V comparable](list *List[V], val V) *List[V] {
	head := &List[V]{}
	head.Next = list
	prev := head
	curr := list
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return head.Next
}

// Compare returns true if the two lists are equal, false otherwise
func Compare[V comparable](list *List[V], other *List[V]) bool {
	for list != nil && other != nil {
		if list.Val != other.Val {
			return false
		}
		list = list.Next
		other = other.Next
	}
	return list == nil && other == nil
}

// Len returns the number of elements in the list
func (l *List[V]) Len() int {
	count := 0
	for l != nil {
		count++
		l = l.Next
	}
	return count
}

// ----------------------------------------------------------------------------
// Utils
// ----------------------------------------------------------------------------

// All returns an iter.Seq[V] that yields all elements in the list
func (l *List[V]) All() iter.Seq[V] {
	return func(yield func(V) bool) {
		for l != nil {
			if !yield(l.Val) {
				return
			}
			l = l.Next
		}
	}
}

// Enumerate returns an iter.Seq2[int, V] that yields the index and value of each element in the list
func (l *List[V]) Enumerate(start int) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := start
		for l != nil {
			if !yield(i, l.Val) {
				return
			}
			i++
			l = l.Next
		}
	}
}

// String returns a string representation of the list
func (l *List[V]) String() string {
	if l == nil {
		return "->[]"
	}
	str := "->["
	for l != nil {
		str += fmt.Sprintf("%v ", l.Val)
		l = l.Next
	}
	if len(str) > 3 {
		str = str[:len(str)-1]
	}
	str += "]"
	return str
}
