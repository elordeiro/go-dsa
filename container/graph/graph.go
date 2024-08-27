package graph

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// Real is a type that emcompasses all signed and unsigned integer types
// as well as all floating point types.
type Real interface {
	constraints.Integer | constraints.Float
}

type Edge[K comparable] struct {
	Src K
	Dst K
}

// Graph is a graph data structure. It is implememted using a hashmap that maps
// vertices to a list of neighbors. The neighbor list itself is also a hashmap
// that maps vertices to an ordered type allowing for the edges to be weighted.
// For unweighted edges, the default weight is set to 1.
type Graph[K comparable, R Real] struct {
	adjList    map[K]map[K]R
	isDirected bool
}

// NewGraph creates a new graph
func NewGraph[K comparable, R Real](isDirected bool) *Graph[K, R] {
	return &Graph[K, R]{adjList: map[K]map[K]R{}, isDirected: isDirected}
}

// AddVertex adds a vertex to the graph
func (g *Graph[K, R]) AddVertex(vertex K) {
	if _, ok := g.adjList[vertex]; !ok {
		g.adjList[vertex] = map[K]R{}
	}
}

// RemoveVertex removes a vertex src the graph. If the vertex doesn't exist, it is a no-op.
func (g *Graph[K, R]) RemoveVertex(vertex K) {
	delete(g.adjList, vertex)
	for _, ns := range g.adjList {
		delete(ns, vertex)
	}
}

// AddEdge adds an edge between two vertices
func (g *Graph[K, R]) AddEdge(src, dst K, weight ...R) {
	var w = R(1)
	if weight != nil {
		w = weight[0]
	}
	g.AddVertex(src)
	g.AddVertex(dst)
	g.adjList[src][dst] = w
	if !g.isDirected {
		g.adjList[dst][src] = w
	}
}

// RemoveEdge removes an edge between two vertices.
func (g *Graph[K, R]) RemoveEdge(src, dst K, weight R) {
	delete(g.adjList[src], dst)
	if !g.isDirected {
		delete(g.adjList[dst], src)
	}
}

// ExistEdge return true if there is an edge between the 2 vertices.
// For undirected graphs, if ExistEdge(src, dst ) returns true, ExistEdge(dst, src)
// will also return true.
func (g *Graph[K, R]) HasEdge(src, dst K) bool {
	if _, ok := g.adjList[src]; ok {
		_, ok := g.adjList[src][dst]
		return ok
	}
	return false
}

// GetNeighbors returns an iter.Se2[K, R] over all the neighbors of a vertex. The first
// returned value is the destination, and the second returned value is the edge weight.
// For unweighted graphs, this value can be ignored since it will always be one.
// Ex:
//
//	Weighted graph -> for v, w := range g.GetNeighbors() { ... }
//	Unweighted graph -> for v := range g.GetNeighbors() { ... }
func (g *Graph[K, R]) Neighbors(vertex K) iter.Seq2[K, R] {
	return func(yield func(K, R) bool) {
		for ns, w := range g.adjList[vertex] {
			if !yield(ns, w) {
				return
			}
		}
	}
}

// Vertices returns an iter.Seq[K] over all vertices in the graph
func (g *Graph[K, R]) Vertices() iter.Seq[K] {
	return func(yield func(K) bool) {
		for v := range g.adjList {
			if !yield(v) {
				return
			}
		}
	}
}

// Edges returns an iter.Seq2[Edge[K], R] over all the edges in the map. The
// first returned value is an Edge[Src, Dst] and the second returned value
// is the edge weight. For unweighted graphs, this value can be ignored since
// it will always be one.
//
//	Ex:
//	Weighted graph:
//	for e, w := range g.Edges() { ... }
//	Unweighted graph:
//	for e := range g.Edges() { ... }
func (g *Graph[K, R]) Edges() iter.Seq2[Edge[K], R] {
	v := make(map[Edge[K]]bool) // visited
	return func(yield func(Edge[K], R) bool) {
		for s, ns := range g.adjList {
			for d, w := range ns {
				e1 := Edge[K]{s, d}
				e2 := Edge[K]{s, d}
				if !v[e1] && !v[e2] {
					if !yield(e1, w) {
						return
					}
				}
			}
		}
	}
}

// Degree returns the out-degree of a vertex
// If the vertex does not exist, this function returns -1.
func (g *Graph[K, R]) Degree(vertex K) int {
	if ns, ok := g.adjList[vertex]; ok {
		return len(ns)
	}
	return -1
}

// Weight returns the weight of an edge between 2 vertices
// If either of the vertices, src or dst, don't exist
// this functin returns 0.
func (g *Graph[K, R]) Weight(src, dst K) R {
	if ns, ok := g.adjList[src]; ok {
		if w, ok := ns[dst]; ok {
			return w
		}
	}
	return 0
}

// SetWeight sets the weight of an edge. If either vertex doesn't exist
// this function does nothing.
func (g *Graph[K, R]) SetWeight(src, dst K, weight R) {
	if ns, ok := g.adjList[src]; ok {
		if _, ok := ns[dst]; ok {
			g.adjList[src][dst] = weight
		}
	}
}

// Has returns true if the graph has the queried vertex and false otherwise.
func (g *Graph[K, R]) Has(vertex K) bool {
	_, ok := g.adjList[vertex]
	return ok
}

// Clone returns a deep copy of the graph
func (g *Graph[K, R]) Clone() *Graph[K, R] {
	clone := NewGraph[K, R](g.isDirected)
	for v, ns := range g.adjList {
		for d, w := range ns {
			clone.AddEdge(v, d, w)
		}
	}
	return clone
}

// Copy copies all the vertices and edges from one graph to another.
// If the destination graph already has a vertex, the edges in the source
// graph will be added to the destination graph.
func (g *Graph[K, R]) Copy(dst *Graph[K, R]) {
	for v, ns := range g.adjList {
		for d, w := range ns {
			dst.AddEdge(v, d, w)
		}
	}
}

// Clear removes all vertices and edges from the graph
func (g *Graph[K, R]) Clear() {
	clear(g.adjList)
}
