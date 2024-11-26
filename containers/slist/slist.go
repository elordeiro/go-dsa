// Package slist provides a singly linked list implementation.
package slist

import (
	"fmt"
	"iter"
	"strings"
)

// List is a singly linked list implementation.
type List[V any] struct {
	Val  V
	Next *List[V]
}

// New creates a new list. If vals are provided the list is initialized with the values.
func New[V any](vals ...V) *List[V] {
	head := &List[V]{}
	curr := head
	for _, val := range vals {
		curr.Next = &List[V]{val, nil}
		curr = curr.Next
	}
	return head.Next
}

// Append adds a new value to the end of the list and returns the new list.
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

// Prepend adds a new value to the beginning of the list and returns the new list.
func (l *List[V]) Prepend(val V) *List[V] {
	return &List[V]{val, l}
}

// Remove removes the first occurrence of the value from the list and returns the new list.
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

// RemoveAll removes all occurrences of the value from the list and returns the new list.
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

// Compare returns true if the two lists are equal, false otherwise.
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

// Len returns the number of elements in the list.
func (l *List[V]) Len() int {
	count := 0
	for l != nil {
		count++
		l = l.Next
	}
	return count
}

// All returns an iter.Seq[V] that yields all elements in the list.
func (l *List[V]) All() iter.Seq[V] {
	return func(yield func(V) bool) {
		head := l
		for l != nil {
			if !yield(l.Val) {
				return
			}
			l = l.Next
		}
		l = head
	}
}

// String returns a string representation of the list.
func (l *List[V]) String() string {
	var sb strings.Builder
	sb.WriteString("->[")
	for l != nil {
		sb.WriteString(fmt.Sprint(l.Val))
		if l.Next != nil {
			sb.WriteByte(' ')
		}
		l = l.Next
	}
	sb.WriteByte(']')
	return sb.String()
}
