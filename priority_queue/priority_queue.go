package pq

import (
	"container/heap"
)

type LessFunc func(item1, item2 any) bool

type PriorityQueue struct {
	items []any
	less  LessFunc
}

func (pq PriorityQueue) Len() int           { return len(pq.items) }
func (pq PriorityQueue) Less(i, j int) bool { return pq.less(pq.items[i], pq.items[j]) }
func (pq PriorityQueue) Swap(i, j int)      { pq.items[i], pq.items[j] = pq.items[j], pq.items[i] }
func (pq *PriorityQueue) Push(x any)        { pq.items = append(pq.items, x) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old.items)
	x := old.items[n-1]
	pq.items = old.items[:n-1]
	return x
}

func NewPriorityQueue(less LessFunc) *PriorityQueue {
	pq := &PriorityQueue{
		items: []any{},
		less:  less,
	}
	heap.Init(pq)
	return pq
}
