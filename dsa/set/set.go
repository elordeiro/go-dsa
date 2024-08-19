package set

import (
	"fmt"
	"iter"
)

type Set[E comparable] map[E]struct{}

func NewSet[E comparable](items ...E) Set[E] {
	s := make(Set[E])
	for _, item := range items {
		s.Add(item)
	}
	return s
}

func (s *Set[E]) Add(item ...E) {
	for _, i := range item {
		(*s)[i] = struct{}{}
	}
}

func (s *Set[E]) Remove(item E) {
	delete(*s, item)
}

func (s *Set[E]) Contains(item E) bool {
	_, found := (*s)[item]
	return found
}

func (s *Set[E]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Set[E]) Len() int {
	return len(*s)
}

func (s *Set[E]) Items() []E {
	items := make([]E, 0, len(*s))
	for item := range *s {
		items = append(items, item)
	}
	return items
}

func (s *Set[E]) String() string {
	items := s.Items()
	return fmt.Sprintf("%v", items)
}

func (s *Set[E]) Union(other *Set[E]) Set[E] {
	union := NewSet(other.Items()...)
	union.Add(s.Items()...)
	return union
}

func (s *Set[E]) Intersection(other *Set[E]) Set[E] {
	intersection := NewSet[E]()
	for item := range *s {
		if other.Contains(item) {
			intersection.Add(item)
		}
	}
	return intersection
}

func (s *Set[E]) Difference(other *Set[E]) Set[E] {
	difference := NewSet[E]()
	for item := range *s {
		if !other.Contains(item) {
			difference.Add(item)
		}
	}
	return difference
}

func (s *Set[E]) Keys() iter.Seq[E] {
	return func(yield func(E) bool) {
		for k := range *s {
			if !yield(k) {
				return
			}
		}
	}
}
