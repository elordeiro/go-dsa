package graph

import (
	"iter"
	"math"
	"slices"

	"github.com/elordeiro/goext/constraints"
	"github.com/elordeiro/goext/containers/deque"
	"github.com/elordeiro/goext/containers/pq"
	"github.com/elordeiro/goext/containers/set"
	"github.com/elordeiro/goext/containers/tuples"
	"github.com/elordeiro/goext/containers/tuples/ituples"
	"github.com/elordeiro/goext/containers/unionfind"
	"github.com/elordeiro/goext/containers/vector"
	extMath "github.com/elordeiro/goext/math"
	"github.com/elordeiro/goext/seqs"
)

// Graph is an interface that defines the methods that a graph must implement
// to be used by the algorithms in this package
type Graph[V comparable, N constraints.Number] interface {
	IsDirected() bool
	Edges() iter.Seq[tuples.Edge[V, N]]
	VertexCount() int
	Vertices() iter.Seq[V]
	Neighbors(V) iter.Seq2[V, N]
}

// option defines the signature of the functions that can be used to configure
// the behavior of the path finding algorithms in this package
type option[V comparable] func(*pathFindOptions[V])

// pathFindOptions holds the configuration of the graph algorithms in this package
type pathFindOptions[V comparable] struct {
	hasBaseCase               bool
	baseCase                  func(V) bool
	vertexFilter, earlyReturn func(V, V) bool
	preVisit                  func(V, V)
	deferred                  func(V)
}

type heuristic[V any, N constraints.Number] func(V) N

// BaseCaseOption returns an option that sets the base case for the path finding
// algorithms. If the predicate returns true, the algorithms will stop and return
// the path.
func BaseCaseOption[V comparable](predicate func(src V) bool) option[V] {
	return func(o *pathFindOptions[V]) {
		o.hasBaseCase = true
		o.baseCase = predicate
	}
}

// VertexFilterOption returns an option that sets the vertex filter for the
// path finding algorithms. If the predicate returns false, the algorithm will
// not visit the vertex.
func VertexFilterOption[V comparable](predicate func(src, dst V) bool) option[V] {
	return func(o *pathFindOptions[V]) {
		o.vertexFilter = predicate
	}
}

// EarlyReturnOption returns an option that sets an early return for the path
// finding algorithms. It is only called on visited vertices. If predicate returns
// true, the algorithm will stop and return the path.
func EarlyReturnOption[V comparable](predicate func(src, dst V) bool) option[V] {
	return func(o *pathFindOptions[V]) {
		o.earlyReturn = predicate
	}
}

// PreVisitOption returns an option that allows for a process to be executed
// before the algorithm visits the vertex. This option cannot modify
// the current path. User may pass in an UnvisitedFilterOption() if they wish
// to modify the path with an unvisited vertex.
func PreVisitOption[V comparable](process func(src, dst V)) option[V] {
	return func(o *pathFindOptions[V]) {
		o.preVisit = process
	}
}

// // PostVisitOption returns an option that allows for a process to be executed
// // after the algorithm visits the vertex. This option cannot modify the current path.
// // User may pass in a VisitedFilterOption() if they wish to modify the path with
// // a visited vertex.
// func PostVisitOption[V comparable](process func(src, dst V)) option[V] {
// 	return func(o *pathFindOptions[V]) {
// 		o.postVisit = process
// 	}
// }

// DeferredOption returns an option that allows for a process to be executed after
// every neighbor of the source vertex has been visited. The process is executed
// only if the source vertex does not satisfy the base case.
func DeferredOption[V comparable](process func(src V)) option[V] {
	return func(o *pathFindOptions[V]) {
		o.deferred = process
	}
}

// EuclideanHeuristic returns a heuristic function that calculates the Euclidean
// distance between two points in a 2D plane.
func EuclideanHeuristic[P ituples.Point2d[N], N constraints.Number](end P) heuristic[P, float64] {
	return func(src P) float64 {
		dx := src.X() - end.X()
		dy := src.Y() - end.Y()
		return math.Sqrt(float64(dx*dx + dy*dy))
	}
}

// setOptions returns a new pathFindOptions with the given options. If no options
// are given, the default options are returned.
func setOptions[V comparable](opts ...option[V]) *pathFindOptions[V] {
	options := &pathFindOptions[V]{
		hasBaseCase:  false,
		baseCase:     func(src V) bool { return false },
		vertexFilter: func(src, dst V) bool { return true },
		earlyReturn:  func(src, dst V) bool { return false },
		preVisit:     func(src, dst V) {},
		deferred:     func(src V) {},
	}
	for _, opt := range opts {
		opt(options)
	}
	return options
}

// Path returns the a path between two vertices. The path is returned as a
// iter.Seq[tuples.Edge[V, N]] If no path exists, an empty sequence is returned.
// The algorithm used is depth-first search.
func Path[V comparable, N constraints.Number](g Graph[V, N], src, dst V) iter.Seq[tuples.Edge[V, N]] {
	baseCase := BaseCaseOption(func(src V) bool { return src == dst })
	return DFS(g, src, baseCase)
}

// HasPath returns true if there is a path between two vertices and false otherwise.
// The algorithm used is depth-first search.
func HasPath[V comparable, N constraints.Number](g Graph[V, N], src, dst V) bool {
	baseCase := BaseCaseOption(func(src V) bool { return src == dst })
	return !seqs.IsEmpty(DFS(g, src, baseCase))
}

// ShortestPath returns the shortest path between two vertices. The path is returned as a
// iter.Seq[tuples.Edges[V, N]]. If no path exists, an empty sequence is returned.
// The algorithm used is breadth-first search. For a faster algorithm, use Dijkstra or AStar.
func ShortesPath[V comparable, N constraints.Number](g Graph[V, N], src, dst V) iter.Seq[tuples.Edge[V, N]] {
	baseCase := BaseCaseOption(func(src V) bool { return src == dst })
	return BFS(g, src, baseCase)
}

// IsConnected returns true if the graph is connected and false otherwise.
// The algorithm used is breadth-first search.
func IsConnected[V comparable, N constraints.Number](g Graph[V, N], start ...V) bool {
	length := g.VertexCount() - 1
	if len(start) > 0 {
		v := start[0]
		return seqs.Len(BFS(g, v)) == length
	}
	for v := range g.Vertices() {
		if seqs.Len(BFS(g, v)) == length {
			return true
		}
	}
	return false
}

// Weight returns the total weight of all the edges in the graph.
func Weight[V comparable, N constraints.Number](g Graph[V, N]) N {
	var total N
	for e := range g.Edges() {
		total += e.Weight()
	}
	return total
}

// HasCycle returns true if the graph has a cycle and false otherwise.
// The algorithm used is depth-first search.
func HasCycle[V comparable, N constraints.Number](g Graph[V, N]) bool {
	v := anyVertex(g)
	var hasCycle bool

	if g.IsDirected() {
		visited := set.New[V]()
		DFS(g, v,
			BaseCaseOption(func(src V) bool {
				visited.Add(src)
				return false
			}),
			EarlyReturnOption(func(src, dst V) bool {
				hasCycle = visited.Contains(dst)
				return hasCycle
			}),
			DeferredOption(func(src V) {
				visited.Remove(src)
			}))
	} else {
		prev := map[V]V{v: v}
		DFS(g, v,
			PreVisitOption(func(src, dst V) {
				prev[dst] = src
			}),
			EarlyReturnOption(func(src, dst V) bool {
				hasCycle = dst != prev[src]
				return hasCycle
			}))
	}

	return hasCycle
}

// DFS performs a depth-first search on the graph g starting from the vertex start.
// The algorithm returns an iter.Seq[tuples.Edge[V, N]] that represents the path from
// the start vertex to the destination vertex which needs to be defined in a base case
// option. If no baseCase is given, DFS will return an empty sequence.
// The algorithm can be configured with the following options:
//
//   - BaseCaseOption: return true if searching should stop, false otherwise.
//   - UnvisitedFilter: return true if unvisited vertex should be considered,
//     false otherwise.
//   - EarlyReturnOption: return true if the search should stop after visiting
//     a certain vertex, false otherwise
//   - PreVisitOption: execute a process with src and dst prior to visiting an unvisited vertex.
//   - PostVisitOpton: execute a process with src and dst after visiting a vertex.
//   - DeferredOption: execute a process with a source vertex if it does not satisfy the base case.
func DFS[V comparable, N constraints.Number](
	g Graph[V, N],
	start V,
	options ...option[V],
) iter.Seq[tuples.Edge[V, N]] {
	opts := setOptions(options...)

	visited := set.New(start)
	path := vector.New[tuples.Edge[V, N]]()

	var dfs func(V) bool
	dfs = func(src V) bool {
		if opts.baseCase(src) {
			return true
		}
		defer opts.deferred(src)
		for dst, w := range g.Neighbors(src) {
			if !visited.Contains(dst) && opts.vertexFilter(src, dst) {
				opts.preVisit(src, dst)
				visited.Add(dst)
				path.Push(tuples.NewEdge(src, dst, w))
				if dfs(dst) {
					return true
				}
				path.Pop()
			} else if opts.earlyReturn(src, dst) {
				return true
			}
		}
		return false
	}
	dfs(start)

	return path.Values()
}

// BFS performs a breadth-first search on the graph g starting from the vertex start.
// The algorithm will return an iter.Seq[tuples.Edge[V, N]] that represents the path
// from the start vertex to the destination vertex which needs to be defined in a base
// case option. If no baseCase is given, BFS will return all of the vertices reachable
// from the starting vertex.
// The algorithm can be configured with the following options:
//
//   - BaseCaseOption: return true if searching should stop, false otherwise.
//   - UnvisitedFilter: return true if unvisited vertex should be considered,
//     false otherwise.
//   - EarlyReturnOption: return true if the search should stop after visiting
//     a certain vertex, false otherwise
//   - PreVisitOption: execute a process with src and dst prior to visiting an unvisited vertex.
//   - PostVisitOpton: execute a process with src and dst after visiting a vertex.
//   - DeferredOption: execute a process with a source vertex if it does not satisfy the base case.
//   - WithDeferred: execute a process with a source vertex if it does not
//     satisfy the base case.
func BFS[V comparable, N constraints.Number](
	g Graph[V, N],
	start V,
	options ...option[V],
) iter.Seq[tuples.Edge[V, N]] {
	opts := setOptions(options...)

	visited := set.New(start)
	q := deque.New(start)
	prev := map[V]tuples.Pair[V, N]{}
	floodFill := vector.New[tuples.Edge[V, N]]()

	var src V
	for !q.IsEmpty() {
		l := q.Len()
		for range l {
			src = q.PopFront()
			if opts.baseCase(src) {
				return rebuildPath(src, start, prev)
			}
			for dst, w := range g.Neighbors(src) {
				if !visited.Contains(dst) && opts.vertexFilter(src, dst) {
					visited.Add(dst)
					opts.preVisit(src, dst)
					q.PushBack(dst)
					floodFill.Push(tuples.NewEdge(src, dst, w))
					prev[dst] = tuples.NewPair(src, w)
				} else if opts.earlyReturn(src, dst) {
					return rebuildPath(src, start, prev)
				}
			}
		}
		opts.deferred(src)
	}

	// no path was found, return either nothing or floodfill
	if opts.hasBaseCase {
		return func(yield func(tuples.Edge[V, N]) bool) {}
	}
	return floodFill.Values()
}

// Dijkstra performs a Dijkstra search on the graph g starting from the vertex start.
// The algorithm will return an iter.Seq[tuples.Edge[V, N]] that represents the path
// from the start vertex to the destination vertex. If no path exists, an empty sequence
// is returned.
// The algorithm can be configured with the following options:
//
//   - BaseCaseOption: return true if searching should stop, false otherwise.
//   - UnvisitedFilter: return true if unvisited vertex should be considered,
//     false otherwise.
//   - EarlyReturnOption: return true if the search should stop after visiting
//     a certain vertex, false otherwise
//   - PreVisitOption: execute a process with src and dst prior to visiting an unvisited vertex.
//   - PostVisitOpton: execute a process with src and dst after visiting a vertex.
//   - DeferredOption: execute a process with a source vertex if it does not satisfy the base case.
func Dijkstra[V comparable, N constraints.Number](
	g Graph[V, N],
	start V,
	options ...option[V],
) iter.Seq[tuples.Edge[V, N]] {
	opts := setOptions(options...)

	pq := pq.NewPQFunc(func(p1, p2 tuples.Pair[V, N]) bool {
		return p1.Right() < p2.Right()
	}, tuples.NewPair(start, N(0)))

	prev := map[V]tuples.Pair[V, N]{}
	dist := map[V]N{}

	inf := extMath.Inf[N]()
	for v := range g.Vertices() {
		dist[v] = inf
	}
	dist[start] = 0

	for !pq.IsEmpty() {
		src := pq.Pop().Left()
		if _, ok := dist[src]; !ok {
			continue
		}
		if opts.baseCase(src) {
			return rebuildPath(src, start, prev)
		}
		for dst, w := range g.Neighbors(src) {
			if _, ok := dist[dst]; ok && opts.vertexFilter(src, dst) {
				if dist[dst] > dist[src]+w {
					opts.preVisit(src, dst)
					dist[dst] = dist[src] + w
					prev[dst] = tuples.NewPair(src, w)
					pq.Push(tuples.NewPair(dst, dist[dst]))
				}
			} else if opts.earlyReturn(src, dst) {
				return rebuildPath(src, start, prev)
			}
		}
		opts.deferred(src)
		delete(dist, src)
	}

	// no path was found
	return func(yield func(tuples.Edge[V, N]) bool) {}
}

// AStar performs an A* search on the graph g starting from the vertex start and
// ending at the vertex end. The algorithm will return an iter.Seq[tuples.Edge[V, N]]
// that represents the path from the start vertex to the destination vertex. If no path
// exists, an empty sequence is returned.
// The algorithm can be configured with the following options:
//
//   - BaseCaseOption: return true if searching should stop, false otherwise.
//   - UnvisitedFilter: return true if unvisited vertex should be considered,
//     false otherwise.
//   - EarlyReturnOption: return true if the search should stop after visiting
//     a certain vertex, false otherwise
//   - PreVisitOption: execute a process with src and dst prior to visiting an unvisited vertex.
//   - PostVisitOpton: execute a process with src and dst after visiting a vertex.
//   - DeferredOption: execute a process with a source vertex if it does not satisfy the base case.
func AStar[V comparable, N constraints.Number](
	g Graph[V, N],
	start V,
	h heuristic[V, N],
	options ...option[V],
) iter.Seq[tuples.Edge[V, N]] {
	opts := setOptions(options...)

	visited := set.New[V]()
	prev := map[V]tuples.Pair[V, N]{}
	gScore := map[V]N{}
	fScore := map[V]N{}

	pq := pq.NewPQFunc(func(p1, p2 tuples.Pair[V, N]) bool {
		return p1.Right() < p2.Right()
	}, tuples.NewPair(start, N(0)))

	f := func(src V) N { return gScore[src] + h(src) }

	inf := extMath.Inf[N]()
	for v := range g.Vertices() {
		fScore[v] = inf
		gScore[v] = inf
	}
	gScore[start] = 0
	fScore[start] = h(start)

	for !pq.IsEmpty() {
		src := pq.Pop().Left()
		if opts.baseCase(src) {
			return rebuildPath(src, start, prev)
		}
		visited.Add(src)
		for dst, w := range g.Neighbors(src) {
			tentative := gScore[src] + w
			if tentative < gScore[dst] {
				prev[dst] = tuples.NewPair(src, w)
				gScore[dst] = tentative
				fScore[dst] = f(dst)
				if !visited.Contains(dst) && opts.vertexFilter(src, dst) {
					opts.preVisit(src, dst)
					pq.Push(tuples.NewPair(dst, fScore[dst]))
				} else if opts.earlyReturn(src, dst) {
					return rebuildPath(src, start, prev)
				}
			}
		}
		opts.deferred(src)
	}

	// Open set is empty but goal was never reached
	return func(yield func(tuples.Edge[V, N]) bool) {}
}

// Kruskal returns the minimum spanning tree of the graph g.
// The algorithm will returns an iter.Seq[tuples.Edge[V, N]] that
// represents the minimum spanning tree of the graph.
func Kruskal[V comparable, N constraints.Number](g Graph[V, N]) iter.Seq[tuples.Edge[V, N]] {
	if g.IsDirected() {
		panic("graph is directed")
	}
	vec := vector.New[tuples.Edge[V, N]]()
	uf := unionfind.New[V]()

	for src := range g.Vertices() {
		uf.MakeSet(src)
		for dst, w := range g.Neighbors(src) {
			vec.Push(tuples.NewEdge(src, dst, w))
		}
	}

	slices.SortFunc(vec, func(a, b tuples.Edge[V, N]) int {
		return int(float64(b.Weight()) - float64(a.Weight()))
	})

	mst := vector.New[tuples.Edge[V, N]]()
	for e := range vec.Backwards() {
		src, dst := e.Src(), e.Dst()
		if !uf.Connected(src, dst) {
			uf.Union(src, dst)
			mst.Push(e)
		}
	}

	return func(yield func(tuples.Edge[V, N]) bool) {
		for e := range mst.Values() {
			if !yield(e) {
				return
			}
		}
	}
}

// Note: Needs more testing
// Edmonds returns the minimum arborescence of the graph g rooted at the vertex root.
func Edmonds[V comparable, N constraints.Number](g Graph[V, N], r V) iter.Seq[tuples.Edge[V, N]] {
	// Since the algorithm requires the creation of new vertices (super nodes)
	// and vertices can be any comparable type, we can rely on new() to create
	// new vertices and use a map to keep track of the original vertices.
	vertexPtrs := map[V]*V{}
	vertices := map[*V]V{}

	vertexPtrs[r] = &r
	vertices[&r] = r
	for e := range BFS(g, r) {
		v := e.Dst()
		vertexPtrs[v] = &v
		vertices[&v] = v
	}

	// collect all the edges that are not self loops and do not point to the root
	edges := []tuples.Edge[*V, N]{}
	for e := range g.Edges() {
		if _, ok := vertexPtrs[e.Src()]; !ok || e.Dst() == r || e.Src() == e.Dst() {
			continue
		}
		u, v, w := vertexPtrs[e.Src()], vertexPtrs[e.Dst()], e.Weight()
		edges = append(edges, tuples.NewEdge(u, v, w))
	}

	// findCycle returns a set of vertices that form a cycle in the graph.
	// The function uses the tortoise and hare algorithm to find the cycle.
	findCycle := func(in map[*V]tuples.Pair[*V, N], start *V) set.Set[*V] {
		slow := start
		fast := in[start].Left()

		for slow != fast {
			slow = in[slow].Left()
			fast = in[in[fast].Left()].Left()
		}

		cycle := set.New[*V]()
		slow = start
		metOnce := false
		for {
			slow = in[slow].Left()
			cycle.Add(fast)
			fast = in[fast].Left()
			cycle.Add(fast)
			fast = in[fast].Left()
			if slow != fast {
				if metOnce {
					break
				}
				metOnce = true
			}
		}
		return cycle
	}

	// contractAndExpand is a recursive function that contracts and expands the graph
	// until a minimum arborescence is found. The function uses the Chu-Liu-Edmonds
	// algorithm to find the minimum arborescence.
	var contractAndExpand func([]tuples.Edge[*V, N]) []tuples.Edge[*V, N]
	contractAndExpand = func(edges []tuples.Edge[*V, N]) []tuples.Edge[*V, N] {
		uf := unionfind.New[*V]()
		in := map[*V]tuples.Pair[*V, N]{} // incoming edge with the minimum weight. v -> (u, w)

		// find the incoming edge with the minimum weight for each vertex
		for _, e := range edges {
			u, v, w := e.Src(), e.Dst(), e.Weight()
			if _, ok := in[v]; !ok {
				uf.MakeSet(v)
				in[v] = tuples.NewPair(u, w)
				continue
			}
			if w < in[v].Right() {
				in[v] = tuples.NewPair(u, w)
			}
		}

		// find a cycle in the graph
		cycle := set.Set[*V]{}
		for v, p := range in {
			u := p.Left()
			if !uf.Connected(u, v) {
				uf.Union(u, v)
			} else {
				// cycle found, but maybe not every vertex with the same representative
				// is part of the cycle
				cycle = findCycle(in, uf.Find(u))
				break
			}
		}

		// if no cycle was found, return the minimum arborescence
		if cycle.Len() == 0 {
			res := []tuples.Edge[*V, N]{}
			for v, p := range in {
				u, w := p.Left(), p.Right()
				res = append(res, tuples.NewEdge(u, v, w))
			}
			return res
		}

		// create a new vertex and edges to contract the cycle
		vc := new(V)                                           // new vertex, the super node
		newEdges := []tuples.Edge[*V, N]{}                     // contracted graph
		mapping := map[tuples.Edge[*V, N]]tuples.Edge[*V, N]{} // map: new edge -> old edge
		for _, e := range edges {
			u, v, w := e.Src(), e.Dst(), e.Weight()
			if !cycle.Contains(u) && cycle.Contains(v) {
				// edge pointing into the cycle
				newE := tuples.NewEdge(u, vc, w-in[v].Right()) // adjust edge weight
				mapping[newE] = e
				newEdges = append(newEdges, newE)
			} else if cycle.Contains(u) && !cycle.Contains(v) {
				// edge pointing out of the cycle
				newE := tuples.NewEdge(vc, v, w)
				mapping[newE] = e
				newEdges = append(newEdges, newE)
			} else if !cycle.Contains(u) && !cycle.Contains(v) {
				// edge not part of the cycle
				newEdges = append(newEdges, e)
			}
		}

		// recursively contract and expand the graph
		newEdges = contractAndExpand(newEdges)

		// vvvvv graph expansion vvvvv

		// remap the edges that point in or out of the cycle
		for i, e := range newEdges {
			u, v := e.Src(), e.Dst()
			if u == vc {
				// edge pointing out of the cycle
				newEdges[i] = mapping[e]
			} else if v == vc {
				// edge pointing into the cycle will replace an edge
				// in the cycle
				newEdges[i] = mapping[e]
				delete(in, mapping[e].Dst())
			}
		}

		// Add the edges that were contracted.
		// At least one edge has been removed from the cycle,
		// so its dst is no longer in 'in' map.
		for m := range cycle.All() {
			if _, ok := in[m]; !ok {
				continue
			}
			u, v, w := in[m].Left(), m, in[m].Right()
			newEdges = append(newEdges, tuples.NewEdge(u, v, w))
		}

		return newEdges
	}

	// find the minimum arborescence
	newEdges := contractAndExpand(edges)

	return func(yield func(tuples.Edge[V, N]) bool) {
		for _, e := range newEdges {
			// convert the new vertices to the original vertices
			u, v, w := vertices[e.Src()], vertices[e.Dst()], e.Weight()
			if !yield(tuples.NewEdge(u, v, w)) {
				return
			}
		}
	}
}

// Needs to implement generic solution using the method used in:
// https://github.com/elordeiro/competitive-programming/blob/main/go/directedmst/main.go
// func Tarjan[V comparable, N constraints.Number](g Graph[V, N], root V) iter.Seq[tuples.Edge[V, N]] {
// }

// rebuildPath returns a function that builds a path from the source vertex to the
// destination vertex. The path is built using the prev map.
func rebuildPath[V comparable, N constraints.Number](
	src V,
	start V,
	prev map[V]tuples.Pair[V, N],
) iter.Seq[tuples.Edge[V, N]] {
	vec := vector.New[tuples.Edge[V, N]]()
	for src != start {
		p := prev[src]
		vec.Push(tuples.NewEdge(p.Left(), src, p.Right()))
		src = p.Left()
	}
	return vec.Backwards()
}

func anyVertex[V comparable, N constraints.Number](g Graph[V, N]) V {
	var rand V
	for v := range g.Vertices() {
		rand = v
		break
	}
	return rand
}
