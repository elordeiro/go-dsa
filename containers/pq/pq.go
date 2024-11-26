// Package pq provides a priority queue implementation that aims to simplify the use of
// priority queues in Go. The package implements a MinPQ, MaxPQ, and a generic PQ that
// can be initialized with a custom less function, as well as a priority queue interface
// that can be used to initialize a priority queue with any container that implements
// pq.Interface.
package pq

import (
	"fmt"
	"iter"

	"github.com/elordeiro/goext/constraints"
	"github.com/elordeiro/goext/containers/heap"
)

// Interface is a priority queue interface that extends the heap.Interface
type Interface[V any] interface {
	heap.Interface[V]
	At(i int) V
	Set(i int, val V)
}

// PQ is a priority queue implementation that uses a heap.Interface to manage the elements
// in the priority queue. The priority queue can be initialized as a min heap, max heap, or
// with a custom less function.
type PQ[V any] struct {
	hp Interface[V]
}

// NewMinPQ creates a new min heap. If vals are provided, they are added to the heap,
// and the heap is initialized.
func NewMinPQ[V constraints.Ordered](vals ...V) *PQ[V] {
	if vals == nil {
		vals = slice[V]{}
	}
	hp := &minHeap[V]{vals}
	heap.Init(hp)
	pq := &PQ[V]{hp}
	return pq
}

// NewMaxPQ creates a new max heap. If vals are provided, they are added to the heap,
// and the heap is initialized.
func NewMaxPQ[V constraints.Ordered](vals ...V) *PQ[V] {
	if vals == nil {
		vals = []V{}
	}
	hp := &maxHeap[V]{vals}
	heap.Init(hp)
	pq := &PQ[V]{hp}
	return pq
}

// NewPQFunc creates a new custom priority queue that can be used for items that are not
// ordered as per the definition of [constraints.Ordered]. If vals are provided, they are
// added to the heap, and the heap is initialized. The less function is used to determine
// the priority of the elements in the priority queue.
func NewPQFunc[V any](less func(V, V) bool, vals ...V) *PQ[V] {
	if vals == nil {
		vals = []V{}
	}
	hp := &funcHeap[V]{vals, less}
	heap.Init(hp)
	pq := &PQ[V]{hp}
	return pq
}

// NewPQFrom creates a new priority queue from an existing PQ.Interface and values. The
// values are added to the heap and the heap is initialized.
func NewPQFrom[V any](hp Interface[V], vals ...V) *PQ[V] {
	if vals == nil {
		vals = []V{}
	}
	for _, val := range vals {
		hp.Push(val)
	}
	heap.Init(hp)
	pq := &PQ[V]{hp}
	return pq
}

// Len returns the number of elements in the priority queue
func (pq *PQ[V]) Len() int {
	return pq.hp.Len()
}

// Push adds an element to the priority queue
func (pq *PQ[V]) Push(val ...V) {
	for _, v := range val {
		heap.Push(pq.hp, v)
	}
}

// Pop removes and returns the element with the highest priority
func (pq *PQ[V]) Pop() V {
	return heap.Pop(pq.hp)
}

// Top returns the element with the highest priority without removing it.
func (pq PQ[V]) Top() V {
	return pq.hp.At(0)
}

// IsEmpty returns true if the priority queue is empty
func (pq *PQ[V]) IsEmpty() bool {
	return pq.Len() == 0
}

// Remove removes and returns the element at index i from the priority queue
func (pq *PQ[V]) Remove(i int) V {
	return heap.Remove(pq.hp, i)
}

// Merge merges the priority queue with another priority queue and returns a new priority queue.
func (pq *PQ[V]) Merge(other *PQ[V]) {
	for v := range other.Drain() {
		pq.Push(v)
	}
}

// Update updates the value at index i in the priority queue and re-establishes
// the heap ordering with a call to Fix. As per the container/heap documentation,
// changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(i) followed by a Push of the new value
func (pq *PQ[V]) Update(i int, val V) {
	pq.hp.Set(i, val)
	heap.Fix(pq.hp, i)
}

// All returns an iter.Seq[V] of values in the priority queue in priority order.
func (pq *PQ[V]) All() iter.Seq[V] {
	return func(yield func(V) bool) {
		elements := make([]V, pq.Len())
		idx := 0

		defer func() {
			for i := range idx {
				pq.Push(elements[i])
			}
		}()

		for !pq.IsEmpty() {
			val := pq.Pop()
			elements[idx] = val
			idx++
			if !yield(val) {
				return
			}
		}
	}
}

// Drain returns an iter.Seq[V] of values in the priority queue in priority order
// by popping all the values from the priority queue. The priority queue is empty
// after calling Drain.
// It returns a single use iterator.
func (pq *PQ[V]) Drain() iter.Seq[V] {
	return func(yield func(V) bool) {
		for !pq.IsEmpty() {
			if !yield(pq.Pop()) {
				return
			}
		}
	}
}

// String returns a string representation of the priority queue
func (pq *PQ[V]) String() string {
	return fmt.Sprint("!", pq.hp)
}

// ----------------------------------------------------------------------------
// Internal heap implementations
// ----------------------------------------------------------------------------

type slice[V any] []V

func (s slice[V]) At(i int) V       { return s[i] }
func (s slice[V]) Len() int         { return len(s) }
func (s slice[V]) Set(i int, val V) { s[i] = val }
func (s slice[V]) Swap(i, j int)    { s[i], s[j] = s[j], s[i] }
func (s *slice[V]) Push(val V)      { *s = append(*s, val) }
func (s *slice[V]) Pop() V {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func (s slice[V]) String() string { return fmt.Sprint([]V(s)) }

type minHeap[V constraints.Ordered] struct {
	slice[V]
}

type maxHeap[V constraints.Ordered] struct {
	slice[V]
}

type funcHeap[V any] struct {
	slice[V]
	less func(V, V) bool
}

func (hp minHeap[V]) Less(i, j int) bool  { return hp.slice[i] < hp.slice[j] }
func (hp maxHeap[V]) Less(i, j int) bool  { return hp.slice[i] > hp.slice[j] }
func (hp funcHeap[V]) Less(i, j int) bool { return hp.less(hp.slice[i], hp.slice[j]) }
