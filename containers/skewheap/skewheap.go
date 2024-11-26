package skewheap

import "github.com/elordeiro/goext/constraints"

type SkewHeap[V any, N constraints.Number] struct {
	Item         V
	offset, cost N
	left, right  *SkewHeap[V, N]
}

func Merge[V any, N constraints.Number](sh1, sh2 *SkewHeap[V, N]) *SkewHeap[V, N] {
	if sh1 == nil {
		return sh2
	}
	if sh2 == nil {
		return sh1
	}
	if sh1.offset != 0 {
		propagate(sh1)
	}
	if sh2.offset != 0 {
		propagate(sh2)
	}
	if sh1.cost > sh2.cost {
		sh1, sh2 = sh2, sh1
	}
	sh1.right = Merge(sh1.right, sh2)
	sh1.left, sh1.right = sh1.right, sh1.left
	return sh1
}

func Push[V any, N constraints.Number](heap *SkewHeap[V, N], cost N, item V) *SkewHeap[V, N] {
	newNode := &SkewHeap[V, N]{Item: item, cost: cost}
	return Merge(heap, newNode)
}

func Pop[V any, N constraints.Number](heap *SkewHeap[V, N]) *SkewHeap[V, N] {
	return Merge(heap.left, heap.right)
}

func Update[V any, N constraints.Number](heap *SkewHeap[V, N], offset N) {
	if heap == nil {
		return
	}
	heap.cost += offset
	heap.offset += offset
}

func propagate[V any, N constraints.Number](sh *SkewHeap[V, N]) {
	if sh.left != nil {
		sh.left.cost += sh.offset
		sh.left.offset += sh.offset
	}
	if sh.right != nil {
		sh.right.cost += sh.offset
		sh.right.offset += sh.offset
	}
	sh.offset = 0
}
