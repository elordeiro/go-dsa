// Package unionfind provides a UnionFind data structure that keeps track of a
// set of elements partitioned into a number of disjoint (non-overlapping) subsets.
package unionfind

import (
	"iter"
	"maps"
	"strings"

	"github.com/elordeiro/goext/containers/set"
)

// UnionFind is a data structure that keeps track of elements partitioned into disjoint sets.
type UnionFind[V comparable] struct {
	parent map[V]V
	rank   map[V]int
}

// New creates a new UnionFind data structure.
func New[V comparable]() *UnionFind[V] {
	return &UnionFind[V]{parent: map[V]V{}, rank: map[V]int{}}
}

// MakeSet creates a new set with a single element.
func (uf *UnionFind[V]) MakeSet(x V) {
	uf.parent[x] = x
	uf.rank[x] = 0
}

// Find returns the representative of the set that contains x.
func (uf *UnionFind[V]) Find(x V) V {
	for x != uf.parent[x] {
		x, uf.parent[x] = uf.parent[x], uf.parent[uf.parent[x]]
	}
	return x
}

// Union merges the sets that contain x and y.
func (uf *UnionFind[V]) Union(x, y V) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return
	}
	if uf.rank[rootX] < uf.rank[rootY] {
		rootX, rootY = rootY, rootX
	}
	uf.parent[rootY] = rootX
	if uf.rank[rootX] == uf.rank[rootY] {
		uf.rank[rootX]++
	}
}

// Connected returns true if x and y are in the same set.
func (uf *UnionFind[V]) Connected(x, y V) bool {
	return uf.Find(x) == uf.Find(y)
}

// Members returns a set of all the elements in the set that contains x.
func (uf *UnionFind[V]) Members(x V) set.Set[V] {
	root := uf.Find(x)
	members := set.New[V]()
	for v := range maps.Keys(uf.parent) {
		if uf.Find(v) == root {
			members.Add(v)
		}
	}
	return members
}

// All returns an iter.Seq[Set[V]] of all the groups in the UnionFind.
func (uf *UnionFind[V]) All() iter.Seq[set.Set[V]] {
	groups := map[V]set.Set[V]{}
	for v := range maps.Keys(uf.parent) {
		root := uf.Find(v)
		if _, ok := groups[root]; !ok {
			groups[root] = set.New[V]()
		}
		groups[root].Add(v)
	}

	return func(yield func(set.Set[V]) bool) {
		for _, group := range groups {
			if !yield(group) {
				return
			}
		}
	}
}

// String returns a string representation of the UnionFind data structure.
func (uf *UnionFind[V]) String() string {
	var sb strings.Builder
	sb.WriteByte('[')
	first := true
	for group := range uf.All() {
		if first {
			first = false
		} else {
			sb.WriteByte(' ')
		}
		sb.WriteByte('~')
		sb.WriteString(group.String())
	}
	sb.WriteByte(']')
	return sb.String()
}
