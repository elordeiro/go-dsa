package pq

import (
	"container/heap"
)

type Heap struct {
	items []any
	less  func(item1, item2 any) bool
}

func (pq Heap) Len() int           { return len(pq.items) }
func (pq Heap) Less(i, j int) bool { return pq.less(pq.items[i], pq.items[j]) }
func (pq Heap) Swap(i, j int)      { pq.items[i], pq.items[j] = pq.items[j], pq.items[i] }
func (pq *Heap) Push(x any)        { pq.items = append(pq.items, x) }
func (pq *Heap) Pop() any {
	old := *pq
	n := len(old.items)
	x := old.items[n-1]
	pq.items = old.items[:n-1]
	return x
}

// Returns a new Max Heap that operates on integers.
// For a Max Heap that operates on any type, use NewHeapFunc.
func NewMaxHeap() *Heap {
	pq := &Heap{
		items: []any{},
		less:  func(item1, item2 any) bool { return item1.(int) < item2.(int) },
	}
	heap.Init(pq)
	return pq
}

// Returns a new Min Heap that operates on integers.
// For a Min Heap that operates on custom types, use NewHeapFunc.
func NewMinHeap() *Heap {
	pq := &Heap{
		items: []any{},
		less:  func(item1, item2 any) bool { return item1.(int) > item2.(int) },
	}
	heap.Init(pq)
	return pq
}

// Returns a new Heap that operates on custom types.
// The less function is used to determine the priority of items.
// If less returns true, item1 is considered to have higher priority than item2.
func NewHeapFunc(less func(item1, item2 any) bool) *Heap {
	pq := &Heap{
		items: []any{},
		less:  less,
	}
	heap.Init(pq)
	return pq
}
