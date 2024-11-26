package graph_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/graph"
	"github.com/elordeiro/goext/containers/hashgraph"
	"github.com/elordeiro/goext/containers/set"
	"github.com/elordeiro/goext/containers/tuples"
	"github.com/elordeiro/goext/seqs"
	"github.com/elordeiro/goext/seqs/transform"
)

func TestHasPath(t *testing.T) {
	g := hashgraph.New[int, int](true)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	tests := []struct {
		start, end int
		want       bool
	}{
		{1, 7, true},
		{1, 5, true},
		{1, 6, true},
		{1, 4, true},
		{1, 8, false},
	}

	for _, tc := range tests {
		got := graph.HasPath(g, tc.start, tc.end)
		if got != tc.want {
			t.Errorf("HasPath(%v, %v) = %v, want %v", tc.start, tc.end, got, tc.want)
		}
	}
}

func TestIsConnected(t *testing.T) {
	g := hashgraph.New[int, int](true)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	if !graph.IsConnected(g) {
		t.Errorf("IsConnected() = false, want true")
	}

	g.AddEdge(8, 9)

	if graph.IsConnected(g) {
		t.Errorf("IsConnected() = true, want false")
	}
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		edges          [][2]int
		directed, want bool
	}{
		{[][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}, true, true},
		{[][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, true, false},
		{[][2]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}}, true, false},
		{[][2]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {4, 3}, {5, 2}, {5, 3}, {5, 4}}, true, false},
		{[][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}, false, true},
		{[][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, false, false},
		{[][2]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {4, 3}, {5, 2}, {5, 3}, {5, 4}}, false, true},
	}

	for _, tc := range tests {
		g := hashgraph.New[int, int](tc.directed)
		for _, e := range tc.edges {
			g.AddEdge(e[0], e[1])
		}

		if got := graph.HasCycle(g); got != tc.want {
			t.Errorf("HasCycle() = %v, want %v", got, tc.want)
		}
	}
}
func TestDfs(t *testing.T) {
	g := hashgraph.New[int, int](true)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	tests := []struct {
		start, end int
		want       [][3]int
	}{
		{1, 7, [][3]int{{1, 3, 1}, {3, 7, 1}}},
		{1, 5, [][3]int{{1, 2, 1}, {2, 5, 1}}},
		{1, 6, [][3]int{{1, 3, 1}, {3, 6, 1}}},
		{1, 4, [][3]int{{1, 2, 1}, {2, 4, 1}}},
		{1, 8, nil},
	}

	for _, tc := range tests {
		want := slices.Values(tuples.Edges(tc.want...))
		got := graph.Path(g, tc.start, tc.end)
		if !seqs.Equal(got, want) {
			t.Errorf("Path(%v, %v) = %v, want %v", tc.start, tc.end, seqs.String(got), seqs.String(want))
		}
	}
}

func TestBfs(t *testing.T) {
	g := hashgraph.New[int, int](true)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	tests := []struct {
		start, end int
		want       [][3]int
	}{
		{1, 7, [][3]int{{1, 3, 1}, {3, 7, 1}}},
		{1, 5, [][3]int{{1, 2, 1}, {2, 5, 1}}},
		{1, 6, [][3]int{{1, 3, 1}, {3, 6, 1}}},
		{1, 4, [][3]int{{1, 2, 1}, {2, 4, 1}}},
		{1, 8, nil},
	}

	for _, tc := range tests {
		want := slices.Values(tuples.Edges(tc.want...))
		got := graph.ShortesPath(g, tc.start, tc.end)
		if !seqs.Equal(got, want) {
			t.Errorf("ShortesPath(%v, %v) = %v, want %v", tc.start, tc.end, seqs.String(got), seqs.String(want))
		}
	}
}

func TestBfs2(t *testing.T) {
	g := hashgraph.New[int, float64](true)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	want := set.New(1, 2, 3)
	got := graph.BFS(g, 0)
	for e := range got {
		if !want.Contains(e.Dst()) {
			t.Errorf("Dfs(g, 0) = %s, want %s", seqs.String(got), seqs.String(want.All()))
			return
		}
	}

	length := seqs.Len(got)
	if length != want.Len() {
		t.Errorf("Dfs() = %v, want %v", length, want.Len())
	}
}

func TestBfsMultiplePathsToDst(t *testing.T) {
	g := hashgraph.New[int, int](true)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(4, 7)
	g.AddEdge(5, 7)

	tests := []struct {
		start, end int
		want       [][3]int
	}{
		{1, 7, [][3]int{{1, 3, 1}, {3, 7, 1}}},
		{1, 5, [][3]int{{1, 2, 1}, {2, 5, 1}}},
		{1, 6, [][3]int{{1, 3, 1}, {3, 6, 1}}},
		{1, 4, [][3]int{{1, 2, 1}, {2, 4, 1}}},
		{1, 8, nil},
	}

	for _, tc := range tests {
		want := slices.Values(tuples.Edges(tc.want...))
		got := graph.ShortesPath(g, tc.start, tc.end)
		if !seqs.Equal(got, want) {
			t.Errorf("ShortesPath(%v, %v) = %v, want %v", tc.start, tc.end, seqs.String(got), seqs.String(want))
		}
	}
}

func TestDijkstra(t *testing.T) {
	tests := []struct {
		start, end int
		edges      [][3]int
		want       [][3]int
	}{
		{
			start: 0,
			end:   4,
			edges: [][3]int{{0, 1, 4}, {0, 2, 1}, {1, 3, 1}, {2, 1, 2}, {2, 3, 5}, {3, 4, 3}},
			want:  [][3]int{{0, 2, 1}, {2, 1, 2}, {1, 3, 1}, {3, 4, 3}},
		},
		{
			start: 1,
			end:   5,
			edges: [][3]int{
				{1, 2, 7}, {1, 3, 9}, {1, 6, 14}, {2, 3, 10}, {2, 4, 15}, {3, 4, 11}, {3, 6, 2}, {4, 5, 6}, {5, 6, 9},
			},
			want: [][3]int{{1, 3, 9}, {3, 6, 2}, {6, 5, 9}},
		},
	}

	for _, tc := range tests {
		g := hashgraph.New[int, int](false)

		edges := slices.Values(tc.edges)
		transform.ForEach(edges, func(e [3]int) {
			g.AddEdge(e[0], e[1], e[2])
		})

		findEnd := graph.BaseCaseOption(func(src int) bool { return src == tc.end })
		want := slices.Values(tuples.Edges(tc.want...))
		got := graph.Dijkstra(g, tc.start, findEnd)

		if !seqs.Equal(got, want) {
			t.Errorf("Dijkstra(g, 0, find%d()) = %s, want %s", tc.end, seqs.String(got), seqs.String(want))
			continue
		}

		gotLen, wantLen := seqs.Len(got), seqs.Len(want)
		if gotLen != wantLen {
			t.Errorf("Dijkstra(g, 0, find%d()) = %s, want %s", tc.end, seqs.String(got), seqs.String(want))
		}
	}
}

func TestAStar(t *testing.T) {
	tests := []struct {
		start, end int
		edges      [][3]int
		want       [][3]int
	}{
		{
			start: 0,
			end:   4,
			edges: [][3]int{{0, 1, 4}, {0, 2, 1}, {1, 3, 1}, {2, 1, 2}, {2, 3, 5}, {3, 4, 3}},
			want:  [][3]int{{0, 2, 1}, {2, 1, 2}, {1, 3, 1}, {3, 4, 3}},
		},
		{
			start: 1,
			end:   5,
			edges: [][3]int{
				{1, 2, 7}, {1, 3, 9}, {1, 6, 14}, {2, 3, 10}, {2, 4, 15}, {3, 4, 11}, {3, 6, 2}, {4, 5, 6}, {5, 6, 9},
			},
			want: [][3]int{{1, 3, 9}, {3, 6, 2}, {6, 5, 9}},
		},
	}

	for _, tc := range tests {
		g := hashgraph.New[int, int](false)

		edges := slices.Values(tc.edges)
		transform.ForEach(edges, func(e [3]int) {
			g.AddEdge(e[0], e[1], e[2])
		})

		heuristic := func(src int) int { return 0 }

		want := slices.Values(tuples.Edges(tc.want...))
		got := graph.AStar(g, tc.start, heuristic,
			graph.BaseCaseOption(func(src int) bool { return src == tc.end }),
		)

		if !seqs.Equal(got, want) {
			t.Errorf("AStar(g, 0, %d, heuristic) = %s, want %s", tc.end, seqs.String(got), seqs.String(want))
		}
	}
}

func TestAStarPoints(t *testing.T) {
	g := hashgraph.New[tuples.Point2d[int], float64](false)

	points := []struct {
		src, dst int
		label    string
	}{
		{2, 15, "a"},
		{10, 19, "b"},
		{8, 17, "c"},
		{12, 9, "d"},
		{15, 14, "e"},
		{16, 20, "f"},
		{28, 17, "z"},
	}

	for _, point := range points {
		p := tuples.NewPoint2d(point.src, point.dst)
		g.AddVertex(p)
		g.AddLabel(p, point.label)
	}

	g.AddEdgeByLabel("a", "b", 35.8)
	g.AddEdgeByLabel("a", "c", 19)
	g.AddEdgeByLabel("b", "e", 84.9)
	g.AddEdgeByLabel("b", "f", 30.4)
	g.AddEdgeByLabel("c", "d", 62.6)
	g.AddEdgeByLabel("c", "e", 76.2)
	g.AddEdgeByLabel("d", "e", 11.7)
	g.AddEdgeByLabel("e", "z", 66.7)
	g.AddEdgeByLabel("f", "z", 197.9)

	var edgesSlice []tuples.Edge[tuples.Point2d[int], float64]
	path := slices.Values([][2]string{{"a", "c"}, {"c", "d"}, {"d", "e"}, {"e", "z"}})
	transform.ForEach(path, func(s [2]string) {
		edge, _ := g.EdgeByLabel(s[0], s[1])
		edgesSlice = append(edgesSlice, edge)
	})

	want := slices.Values(edgesSlice)
	start, _ := g.VertexByLabel("a")
	end, _ := g.VertexByLabel("z")
	got := graph.AStar(
		g,
		start,
		graph.EuclideanHeuristic(end),
		graph.BaseCaseOption(func(src tuples.Point2d[int]) bool { return src == end }),
	)

	if !seqs.Equal(got, want) {
		t.Errorf("\nAStar() = \n\t%s, \nwant \n\t%s", seqs.String(got), seqs.String(want))
	}
}

func TestKruskal(t *testing.T) {
	g := hashgraph.New[string, int](false)
	g.AddEdge("a", "b", 1)
	g.AddEdge("a", "d", 4)
	g.AddEdge("a", "e", 3)
	g.AddEdge("b", "d", 4)
	g.AddEdge("b", "e", 2)
	g.AddEdge("c", "e", 4)
	g.AddEdge("c", "f", 5)
	g.AddEdge("d", "e", 4)
	g.AddEdge("e", "f", 7)

	mst := graph.Kruskal(g)
	got := transform.Reduce(mst,
		func(acc int, e tuples.Edge[string, int]) int { return acc + e.Weight() },
	)

	if got != 16 {
		t.Errorf("Weight() = %v, want 16", got)
	}

	g2 := hashgraph.New[string, int](false)
	transform.ForEach(mst, func(e tuples.Edge[string, int]) {
		g2.AddEdge(e.Src(), e.Dst(), e.Weight())
	})

	if graph.HasCycle(g2) {
		t.Errorf("HasCycle() = true, want false")
	}
}

// Waiting on implementation of Tarjan function.
// func TestTarjan(t *testing.T) {
// 	testOnly := 0
// 	tests := []struct {
// 		edges  [][3]int
// 		want   []set.Set[tuples.Edge[int, int]]
// 		weight int
// 		name   string
// 	}{
// 		{
// 			[][3]int{{1, 2, 1}, {1, 3, 4}, {2, 3, 2}, {2, 4, 5}, {3, 4, 1}},
// 			[]set.Set[tuples.Edge[int, int]]{set.New(tuples.Edges([][3]int{{1, 2, 1}, {2, 3, 2}, {3, 4, 1}}...)...)},
// 			4,
// 			"simple graph",
// 		},
// 		{
// 			[][3]int{{1, 2, 2}, {2, 3, 2}, {3, 4, 2}, {4, 2, 1}, {1, 4, 10}},
// 			[]set.Set[tuples.Edge[int, int]]{set.New(tuples.Edges([][3]int{{1, 2, 2}, {2, 3, 2}, {3, 4, 2}}...)...)},
// 			6,
// 			"single cycle",
// 		},
// 		{
// 			[][3]int{{1, 2, 2}, {2, 3, 3}, {3, 4, 4}, {4, 2, 1}, {1, 4, 10}},
// 			[]set.Set[tuples.Edge[int, int]]{set.New(tuples.Edges([][3]int{{1, 2, 2}, {2, 3, 3}, {3, 4, 4}}...)...)},
// 			9,
// 			"graph with cycle",
// 		},
// 		{
// 			[][3]int{{1, 2, 3}, {1, 3, 1}, {2, 3, 4}, {4, 2, 2}, {3, 4, 5}},
// 			[]set.Set[tuples.Edge[int, int]]{set.New(tuples.Edges([][3]int{{1, 3, 1}, {4, 2, 2}, {3, 4, 5}}...)...)},
// 			8,
// 			"multiple incoming edges",
// 		},
// 		{
// 			[][3]int{{1, 2, 1}, {2, 3, 2}, {4, 5, 3}},
// 			[]set.Set[tuples.Edge[int, int]]{set.New(tuples.Edges([][3]int{{1, 2, 1}, {2, 3, 2}}...)...)},
// 			3,
// 			"disconnected graph",
// 		},
// 		{
// 			[][3]int{{1, 2, 1}, {1, 3, 1}, {2, 4, 2}, {3, 4, 2}},
// 			[]set.Set[tuples.Edge[int, int]]{
// 				set.New(tuples.Edges([][3]int{{1, 2, 1}, {2, 4, 2}, {1, 3, 1}}...)...),
// 				set.New(tuples.Edges([][3]int{{1, 3, 1}, {3, 4, 2}, {1, 2, 1}}...)...),
// 			},
// 			4,
// 			"multiple solutions",
// 		},
// 		{
// 			[][3]int{},
// 			[]set.Set[tuples.Edge[int, int]]{},
// 			0,
// 			"single node",
// 		},
// 		{
// 			[][3]int{{1, 2, 3}, {2, 2, 0}, {2, 3, 2}},
// 			[]set.Set[tuples.Edge[int, int]]{set.New(tuples.Edges([][3]int{{1, 2, 3}, {2, 3, 2}}...)...)},
// 			5,
// 			"self loop",
// 		},
// 		{
// 			[][3]int{{1, 2, -5}, {2, 3, -2}, {3, 1, -1}, {3, 4, 3}},
// 			[]set.Set[tuples.Edge[int, int]]{set.New(tuples.Edges([][3]int{{1, 2, -5}, {2, 3, -2}, {3, 4, 3}}...)...)},
// 			-4,
// 			"negative weights",
// 		},
// 		{
// 			[][3]int{{1, 2, 1}, {2, 3, 2}, {3, 4, 3}, {4, 2, 1}, {5, 4, 5}, {3, 5, 4}},
// 			[]set.Set[tuples.Edge[int, int]]{set.New(
// 				tuples.Edges([][3]int{{1, 2, 1}, {2, 3, 2}, {3, 4, 3}, {3, 5, 4}}...,
// 				)...)},
// 			10,
// 			"overlapping cycles",
// 		},
// 		{
// 			[][3]int{{1, 2, 3}, {2, 3, 5}, {3, 4, 2}, {4, 5, 7}, {5, 2, 1}, {5, 6, 10}, {6, 1, 6}, {4, 6, 8}},
// 			[]set.Set[tuples.Edge[int, int]]{set.New(
// 				tuples.Edges([][3]int{{1, 2, 3}, {2, 3, 5}, {3, 4, 2}, {4, 5, 7}, {4, 6, 8}}...,
// 				)...)},
// 			25,
// 			"nested cycles",
// 		},
// 		{
// 			[][3]int{{1, 2, 10}, {1, 3, 12}, {2, 4, 15}, {3, 5, 10}, {4, 2, 5}, {5, 3, 8}, {5, 6, 3}, {4, 6, 7}},
// 			[]set.Set[tuples.Edge[int, int]]{
// 				set.New(tuples.Edges([][3]int{{1, 2, 10}, {1, 3, 12}, {2, 4, 15}, {3, 5, 10}, {5, 6, 3}}...)...)},
// 			50,
// 			"large graph",
// 		},
// 		{
// 			[][3]int{{1, 2, 1}, {2, 3, 2}, {3, 4, 3}, {4, 2, 1}, {5, 6, 4}, {6, 7, 5}, {7, 5, 3}, {1, 5, 6}},
// 			[]set.Set[tuples.Edge[int, int]]{set.New(
// 				tuples.Edges([][3]int{{1, 2, 1}, {2, 3, 2}, {3, 4, 3}, {1, 5, 6}, {5, 6, 4}, {6, 7, 5}}...,
// 				)...)},
// 			21,
// 			"multiple cycles",
// 		},
// 	}

// 	for i, tc := range tests {
// 		if testOnly > 0 && testOnly != i+1 {
// 			continue
// 		}
// 		g := hashgraph.New[int, int](true)
// 		for _, edge := range tc.edges {
// 			g.AddEdge(edge[0], edge[1], edge[2])
// 		}
// 		if len(tc.edges) == 0 {
// 			g.AddVertex(1)
// 		}

// 		t.Run(tc.name, func(t *testing.T) {
// 			mat := graph.Tarjan(g, 1)
// 			got := set.New[tuples.Edge[int, int]]()
// 			for e := range mat {
// 				got.Add(e)
// 			}

// 			if got.Len() == 0 && len(tc.want) == 0 {
// 				return
// 			}

// 			for _, s := range tc.want {
// 				diff1 := s.Difference(got)
// 				diff2 := got.Difference(s)
// 				if diff1.Len() == 0 && diff2.Len() == 0 {
// 					weightTotal := 0
// 					for edge := range got.All() {
// 						weightTotal += edge.Weight()
// 					}
// 					if weightTotal == tc.weight {
// 						return
// 					}
// 				}
// 			}

// 			t.Errorf("Edmonds(%v, 1) = \n%v, \nwant \n%v", g, got, tc.want)
// 		})
// 	}
// }
