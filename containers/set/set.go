// Package set provides a set data structure that can be used to store unique elements.
// The set is implemented using a map[K]struct{} where K is a comparable type.
package set

import (
	"fmt"
	"iter"
	"slices"
)

// Set is a collection of unique elements.
type Set[K comparable] map[K]struct{}

// New creates a new set with the provided items.
func New[K comparable](items ...K) Set[K] {
	s := make(Set[K])
	for _, item := range items {
		s.Add(item)
	}
	return s
}

// Len returns the number of elements in the set.
func (s Set[K]) Len() int {
	return len(s)
}

// Add adds the provided items to the set.
func (s Set[K]) Add(item ...K) {
	for _, i := range item {
		s[i] = struct{}{}
	}
}

// Remove removes the provided item from the set.
func (s Set[K]) Remove(item K) {
	delete(s, item)
}

// Contains returns true if the provided item is in the set.
func (s Set[K]) Contains(item K) bool {
	_, ok := s[item]
	return ok
}

// IsEmpty returns true if the set is empty.
func (s Set[K]) IsEmpty() bool {
	return len(s) == 0
}

// Equal returns true if the provided set is equal to the set.
// Two sets are equal if they have the same elements.
func (s Set[K]) Equal(other Set[K]) bool {
	if len(s) != len(other) {
		return false
	}
	for item := range s {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// Union returns a new set with the union of the provided set and the set.
func (s Set[K]) Union(other Set[K]) Set[K] {
	union := New(slices.Collect(other.All())...)
	union.Add(slices.Collect(s.All())...)
	return union
}

// Intersection returns a new set with the intersection of the provided set and the set.
func (s Set[K]) Intersection(other Set[K]) Set[K] {
	intersection := New[K]()
	for item := range s {
		if other.Contains(item) {
			intersection.Add(item)
		}
	}
	return intersection
}

// Difference returns a new set with the difference of the provided set and the set.
func (s Set[K]) Difference(other Set[K]) Set[K] {
	difference := New[K]()
	for item := range s {
		if !other.Contains(item) {
			difference.Add(item)
		}
	}
	return difference
}

// All returns an iter.Seq[K] of all elements in the set.
func (s Set[K]) All() iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range s {
			if !yield(k) {
				return
			}
		}
	}
}

// String returns a string representation of the set.
func (s Set[K]) String() string {
	str := []byte(fmt.Sprint(slices.Collect(s.All())))
	str[0], str[len(str)-1] = '{', '}'
	return string(str)
}
