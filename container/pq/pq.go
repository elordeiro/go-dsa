package pq

import (
	"container/heap"
	"iter"

	"golang.org/x/exp/constraints"
)

type Ord constraints.Ordered

// ----------------------------------------------------------------------------
// Internal heap implementation
// ----------------------------------------------------------------------------

type pq[V any] struct {
	Items []V
	less  func(V, V) bool
}

func (pq pq[V]) Len() int           { return len(pq.Items) }
func (pq pq[V]) Less(i, j int) bool { return pq.less(pq.Items[i], pq.Items[j]) }
func (pq pq[V]) Swap(i, j int)      { pq.Items[i], pq.Items[j] = pq.Items[j], pq.Items[i] }
func (pq *pq[V]) Push(val any)      { pq.Items = append(pq.Items, val.(V)) }
func (pq *pq[V]) Pop() any {
	old := *pq
	n := len(old.Items)
	x := old.Items[n-1]
	pq.Items = old.Items[:n-1]
	return x
}

// ----------------------------------------------------------------------------
// Priority Queue
// ----------------------------------------------------------------------------

// Pq is a priority queue implementation
type Pq[V any] struct {
	Internal *pq[V]
}

// ----------------------------------------------------------------------------
// Public API
// ----------------------------------------------------------------------------

// NewMaxHeap creates a new max heap. If vals are provided, they are added to the heap,
// and the heap is initialized. The heap is initialized with the default less function
func NewMaxHeap[V Ord](vals ...V) *Pq[V] {
	if vals == nil {
		vals = []V{}
	}
	pq := &pq[V]{
		Items: vals,
		less:  func(item1, item2 V) bool { return item1 > item2 },
	}
	heap.Init(pq)

	return &Pq[V]{pq}
}

// NewMinHeap creates a new min heap. If vals are provided, they are added to the heap,
// and the heap is initialized. The heap is initialized with the default less function
func NewMinHeap[V Ord](vals ...V) *Pq[V] {
	if vals == nil {
		vals = []V{}
	}
	pq := &pq[V]{
		Items: vals,
		less:  func(item1, item2 V) bool { return item1 < item2 },
	}
	heap.Init(pq)

	return &Pq[V]{pq}
}

// NewPqFunc creates a new priority queue. If vals are provided, they are added to the heap,
// and the heap is initialized. The heap is initialized with the provided less function
func NewPqFunc[V any](less func(item1, item2 V) bool, vals ...V) *Pq[V] {
	if vals == nil {
		vals = []V{}
	}
	pq := &pq[V]{
		Items: vals,
		less:  less,
	}
	heap.Init(pq)

	return &Pq[V]{pq}
}

// Len returns the number of elements in the priority queue
func (pq *Pq[V]) Len() int {
	return pq.Internal.Len()
}

// Push adds an element to the priority queue
func (pq *Pq[V]) Push(val V) {
	heap.Push(pq.Internal, val)
}

// Pop removes and returns the element with the highest priority
func (pq *Pq[V]) Pop() V {
	return heap.Pop(pq.Internal).(V)
}

// Peek returns the element with the highest priority without removing it.
// If the priority queue is empty, this function panics. Use IsEmpty to check if the
// priority queue is empty before calling Peek
func (pq *Pq[V]) Peek() V {
	if pq.IsEmpty() {
		panic("Pq is empty")
	}
	return pq.Internal.Items[0]
}

// IsEmpty returns true if the priority queue is empty
func (pq *Pq[V]) IsEmpty() bool {
	return pq.Len() == 0
}

// Values returns an iter.Seq[V] of values in the priority queue in priority order
// that can be used to iterate over the values in the priority queue
func (pq *Pq[V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for !pq.IsEmpty() {
			if !yield(pq.Pop()) {
				return
			}
		}
	}
}

// All returns an iter.Seq2[int, V] of values in the priority queue in priority order
// that can be used to iterate over the values in the priority queue with their index
func (pq *Pq[V]) All() iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := 0
		for !pq.IsEmpty() {
			if !yield(i, pq.Pop()) {
				return
			}
			i++
		}
	}
}

// Remove removes and returns the element at index i from the priority queue
func (pq *Pq[V]) Remove(i int) V {
	return heap.Remove(pq.Internal, i).(V)
}

// Clear removes all elements from the priority queue
func (pq *Pq[V]) Clear() {
	pq.Internal.Items = []V{}
}

// Update updates the value at index i in the priority queue and re-establishes
// the heap ordering with a call to Fix. As per the container/heap documentation,
// changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(i) followed by a Push of the new value
func (pq *Pq[V]) Update(i int, val V) {
	pq.Internal.Items[i] = val
	heap.Fix(pq.Internal, i)
}

// UpdateLess updates the less function used to determine the priority of elements in the priority queue
// and re-establishes the heap ordering. This is useful when the priority of elements in the priority queue
// is determined by a function that changes over time
func (pq *Pq[V]) UpdateLess(less func(item1, item2 V) bool) {
	pq.Internal.less = less
	heap.Init(pq.Internal)
}

// UpdateAll updates all values in the priority queue and re-establishes the heap ordering
func (pq *Pq[V]) UpdateAll(vals ...V) {
	pq.Internal.Items = vals
	heap.Init(pq.Internal)
}
