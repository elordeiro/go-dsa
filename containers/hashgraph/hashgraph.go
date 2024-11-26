// Package hashgraph provides a graph data structure implemented using a hashmap.
package hashgraph

import (
	"fmt"
	"iter"
	"strings"

	"github.com/elordeiro/goext/constraints"
	"github.com/elordeiro/goext/containers/tuples"
)

// VertexError is returned when a vertex is not found in the graph.
type VertexError[V comparable] struct {
	vertex V
}

func (e VertexError[V]) Error() string {
	return fmt.Sprintf("vertex not found: %v", e.vertex)
}

// EdgeError is returned when an edge is not found in the graph.
type EdgeError[V comparable, N constraints.Number] struct {
	src, dst V
}

func (e EdgeError[V, N]) Error() string {
	return fmt.Sprintf("edge not found: %v -> %v", e.src, e.dst)
}

// HashGraph is a graph data structure. It is implememted using a hashmap that maps
// vertices to a list of neighbors. The neighbor list itself is also a hashmap
// that maps vertices to an number type allowing for the edges to be weighted.
// For unweighted edges, the default weight is set to 1. The graph can be directed
// or undirected. The graph also supports labels for vertices.
type HashGraph[V comparable, N constraints.Number] struct {
	adjList    map[V]map[V]N
	labels     map[string]V
	isDirected bool
}

// New creates a new graph with the specified directedness.
func New[V comparable, N constraints.Number](isDirected bool) *HashGraph[V, N] {
	return &HashGraph[V, N]{adjList: map[V]map[V]N{}, isDirected: isDirected, labels: map[string]V{}}
}

// AddVertex adds a vertex to the graph. If the vertex already exists, it is a no-op.
func (g *HashGraph[V, N]) AddVertex(vertex V) {
	if _, ok := g.adjList[vertex]; !ok {
		g.adjList[vertex] = map[V]N{}
	}
}

// RemoveVertex removes a vertex from the graph. If the vertex doesn't exist, it is a no-op.
func (g *HashGraph[V, N]) RemoveVertex(vertex V) {
	delete(g.adjList, vertex)
	for _, ns := range g.adjList {
		delete(ns, vertex)
	}
}

// HasVertex returns true if the graph has the queried vertex and false otherwise.
func (g HashGraph[V, N]) HasVertex(vertex V) bool {
	_, ok := g.adjList[vertex]
	return ok
}

// Degree returns the out-degree of a vertex in a directed graph and the degree
// of a vertex in an undirected graph. If the vertex doesn't exist, it returns
// a VertexError.
func (g HashGraph[V, N]) Degree(vertex V) (int, error) {
	if ns, ok := g.adjList[vertex]; ok {
		return len(ns), nil
	}
	return 0, VertexError[V]{vertex: vertex}
}

// AddEdge adds an edge between two vertices, src and dst. If the either vertex
// doesn't exist, it is added to the graph. If the graph is undirected, the edge
// is added in both directions. If the edge already exists, the weight is updated.
// If no weight is provided, the default weight is set to 1.
func (g *HashGraph[V, N]) AddEdge(src, dst V, weight ...N) {
	var w N = 1
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

// RemoveEdge removes an edge between two vertices. If the edge doesn't exist, it is a no-op.
func (g *HashGraph[V, N]) RemoveEdge(src, dst V) {
	delete(g.adjList[src], dst)
	if !g.isDirected {
		delete(g.adjList[dst], src)
	}
}

// HasEdge returns true if there is an edge between the 2 vertices.
// For undirected graphs, if HasEdge(src, dst) returns true, HasEdge(dst, src)
// will also return true.
func (g HashGraph[V, N]) HasEdge(src, dst V, weight ...N) bool {
	if _, ok := g.adjList[src]; ok {
		_, ok := g.adjList[src][dst]
		if len(weight) == 0 {
			return ok
		}
		return ok && g.adjList[src][dst] == weight[0]
	}
	return false
}

// Edge returns the edge between two vertices. If the edge doesn't exist, it returns
// an EdgeError.
func (g HashGraph[V, N]) Edge(src, dst V) (tuples.Edge[V, N], error) {
	if ns, ok := g.adjList[src]; ok {
		if w, ok := ns[dst]; ok {
			return tuples.NewEdge(src, dst, w), nil
		}
	}
	return tuples.Edge[V, N]{}, EdgeError[V, N]{src: src, dst: dst}
}

// EdgeWeight returns the weight of an edge between 2 vertices
// If either of the vertices, src or dst, don't exist
// this functin returns 0.
func (g HashGraph[V, N]) EdgeWeight(src, dst V) (N, error) {
	if ns, ok := g.adjList[src]; ok {
		if w, ok := ns[dst]; ok {
			return w, nil
		}
	}
	return 0, EdgeError[V, N]{src: src, dst: dst}
}

// SetEdgeWeight sets the weight of an edge. If either vertex doesn't exist
// it is a no-op.
func (g *HashGraph[V, N]) SetEdgeWeight(src, dst V, weight N) {
	if ns, ok := g.adjList[src]; ok {
		if _, ok := ns[dst]; ok {
			g.adjList[src][dst] = weight
		}
	}
}

// AddLabel adds a label to a vertex. If the label already exists, it is updated.
func (g *HashGraph[V, N]) AddLabel(vertex V, label string) {
	g.labels[label] = vertex
}

// AddEdgeByLabel adds an edge between two vertices using their labels. If either
// vertex doesn't exist, it returns a VertexError.
func (g *HashGraph[V, N]) AddEdgeByLabel(src, dst string, weight ...N) error {
	v1, ok := g.labels[src]
	if !ok {
		return VertexError[string]{vertex: src}
	}
	v2, ok := g.labels[dst]
	if !ok {
		return VertexError[string]{vertex: dst}
	}
	g.AddEdge(v1, v2, weight...)
	return nil
}

// EdgeByLabel returns the edge between two vertices using their labels. If either
// vertex doesn't exist, it returns a VertexError.
func (g HashGraph[V, N]) EdgeByLabel(src, dst string) (tuples.Edge[V, N], error) {
	v1, ok := g.labels[src]
	if !ok {
		return tuples.Edge[V, N]{}, VertexError[string]{vertex: src}
	}
	v2, ok := g.labels[dst]
	if !ok {
		return tuples.Edge[V, N]{}, VertexError[string]{vertex: dst}
	}

	edge := tuples.NewEdge(v1, v2, g.adjList[v1][v2])
	return edge, nil
}

// VertexByLabel returns the vertex with the given label. If the vertex doesn't exist,
// it returns a VertexError.
func (g HashGraph[V, N]) VertexByLabel(label string) (V, error) {
	v, ok := g.labels[label]
	if !ok {
		return v, VertexError[string]{vertex: label}
	}
	return v, nil
}

// VertexCount returns the number of vertices in the graph.
func (g HashGraph[V, N]) VertexCount() int {
	return len(g.adjList)
}

// Clear removes all vertices and edges from the graph
func (g *HashGraph[V, N]) Clear() {
	clear(g.adjList)
	clear(g.labels)
}

// Clone returns a deep copy of the graph.
func (g HashGraph[V, N]) Clone() *HashGraph[V, N] {
	clone := New[V, N](g.isDirected)
	for v, ns := range g.adjList {
		for d, w := range ns {
			clone.AddEdge(v, d, w)
		}
	}
	return clone
}

// Transpose returns the transpose of the graph. The transpose of a graph is a
// graph with all the edges reversed. For example, if there is an edge from
// vertex A to vertex B in the original graph, there will be an edge from
// vertex B to vertex A in the transpose graph. If the graph is undirected,
// the transpose will be a clone of the original graph.
func (g HashGraph[V, N]) Transpose() *HashGraph[V, N] {
	if !g.IsDirected() {
		return g.Clone()
	}
	newGragh := New[V, N](true)
	for src := range g.Vertices() {
		for dst, w := range g.Neighbors(src) {
			newGragh.AddEdge(dst, src, w)
		}
	}
	return newGragh
}

// IsDirected returns true if the graph is directed and false otherwise.
func (g HashGraph[V, N]) IsDirected() bool {
	return g.isDirected
}

// Edges returns an iter.Seq[tuples.Edge[V, N]] over all the edges in the map.
func (g HashGraph[V, N]) Edges() iter.Seq[tuples.Edge[V, N]] {
	return func(yield func(tuples.Edge[V, N]) bool) {
		v := make(map[tuples.Edge[V, N]]bool) // visited
		for s, ns := range g.adjList {
			for d, w := range ns {
				e1 := tuples.NewEdge(s, d, w)
				e2 := e1
				if !g.isDirected {
					e2 = tuples.NewEdge(d, s, w)
				}
				if !v[e1] && !v[e2] {
					v[e1], v[e2] = true, true
					if !yield(e1) {
						return
					}
				}
			}
		}
	}
}

// Vertices returns an iter.Seq[V] over all vertices in the graph.
func (g HashGraph[V, N]) Vertices() iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range g.adjList {
			if !yield(v) {
				return
			}
		}
	}
}

// GetNeighbors returns an iter.Seq2[V, N] over all the neighbors of a vertex. The first
// returned value is the destination, and the second returned value is the edge weight.
// For unweighted graphs, this value can be ignored since it will always be one.
// Ex:
//
//	Weighted graph -> for v, w := range g.GetNeighbors() { ... }
//	Unweighted graph -> for v := range g.GetNeighbors() { ... }
func (g HashGraph[V, N]) Neighbors(vertex V) iter.Seq2[V, N] {
	return func(yield func(V, N) bool) {
		for ns, w := range g.adjList[vertex] {
			if !yield(ns, w) {
				return
			}
		}
	}
}

// String returns a string representation of the graph.
func (g HashGraph[V, N]) String() string {
	var sb strings.Builder
	sb.WriteString("G[")
	first := true
	for e := range g.Edges() {
		if first {
			first = false
		} else {
			sb.WriteByte(' ')
		}
		sb.WriteString(e.String())
	}
	sb.WriteString("]")
	return sb.String()
}
